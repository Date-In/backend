package main

import (
	"context"
	"dating_service/configs"
	"dating_service/internal/action"
	"dating_service/internal/activity"
	"dating_service/internal/auth"
	"dating_service/internal/cache"
	"dating_service/internal/chat"
	"dating_service/internal/dictionaries"
	"dating_service/internal/filestorage"
	"dating_service/internal/filter"
	"dating_service/internal/like"
	"dating_service/internal/match"
	"dating_service/internal/message"
	"dating_service/internal/notifier"
	"dating_service/internal/photo"
	"dating_service/internal/profile"
	"dating_service/internal/recommendations"
	"dating_service/internal/user"
	"dating_service/pkg/JWT"
	"dating_service/pkg/cryptohelper"
	db2 "dating_service/pkg/db"
	"dating_service/pkg/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"time"
)

// @Title API для Сервиса Знакомств
// @Version 1.0
// @Description Это серверная часть для приложения знакомств. Все эндпоинты, требующие авторизации, ожидают JWT токен в заголовке 'Authorization: Bearer {token}'.
// @TermsOfServiceUrl http://swagger.io/terms/

// @ContactName Ваше Имя
// @ContactEmail ваш.email@example.com

// @LicenseName Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html

// @Server http://localhost:8081 Локальный сервер для разработки
// @Server http://185.61.254.35:8081 Документация

// @Security AuthorizationHeader read write
// @SecurityScheme AuthorizationHeader http bearer Input your token
func main() {
	now := time.Now()
	config := configs.NewConfig()
	tokenGenerator := JWT.NewJWT(config.SecretToken.Token)
	db := db2.NewDb(config)
	photoS3Client := filestorage.NewS3Client(config)
	ctx := context.Background()

	//routers
	mainRouter := http.NewServeMux()
	publicRouter := http.NewServeMux()
	protectedRouter := http.NewServeMux()
	//repository
	userRepository := user.NewRepository(db)
	photoRepository := photo.NewRepository(db)
	filterRepository := filter.NewRepository(db)
	actionsRepository := action.NewRepository(db)
	likeRepository := like.NewRepository(db)
	matchRepository := match.NewRepository(db)
	messageRepository := message.NewRepository(db)
	activityRepository := activity.NewRepository(db)
	refCacheRepository := cache.NewRepository(db)
	//ws-hub
	notifierHub := notifier.NewHub()
	go notifierHub.Run()
	//cache
	refCache, err := cache.NewReferenceCache(refCacheRepository)
	if err != nil {
		panic(err)
	}
	//service
	cryptoMessage := cryptohelper.NewService(config)
	dictionariesService := dictionaries.NewService(refCache)
	photoS3Service := filestorage.NewS3FileStorage(photoS3Client, config)
	photoService := photo.NewService(photoRepository, photoS3Service)
	userService := user.NewService(userRepository)
	authService := auth.NewService(userService, refCache, tokenGenerator)
	profileService := profile.NewService(userService, photoService, refCache)
	filterService := filter.NewService(filterRepository)
	recommendationService := recommendations.NewService(userService, filterService)
	actionService := action.NewService(userService, actionsRepository)
	matchService := match.NewService(matchRepository)
	likeService := like.NewService(likeRepository, userService, matchService)
	activityService := activity.NewService(activityRepository)
	messageService := message.NewService(messageRepository, cryptoMessage)
	notifierService := notifier.NewService(notifierHub, activityService, matchService)
	chatService := chat.NewService(matchService, messageService, notifierService)
	//background tasks
	//set non active status
	go func() {
		ticker := time.NewTicker(24 * time.Hour)
		defer ticker.Stop()
		for {
			actionService.ChangeStatusToNonActive()
			<-ticker.C
		}
	}()
	//delete match
	go func() {
		ticker := time.NewTicker(24 * time.Hour)
		defer ticker.Stop()
		for {
			err = matchService.CleanupInactiveMatch()
			if err != nil {
				log.Println(err)
			}
			<-ticker.C
		}
	}()
	//handler-public
	fileServer := http.FileServer(http.Dir("./"))
	publicRouter.Handle("/swagger/oas.yaml", http.StripPrefix("/swagger/", fileServer))
	publicRouter.HandleFunc("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/oas.yaml"),
	))
	dictionaries.NewHandler(publicRouter, dictionariesService)
	auth.NewHandler(publicRouter, authService)
	notifier.NewHandlerWs(publicRouter, notifierService, config)
	chat.NewHandlerWs(publicRouter, chatService, config)
	//handler-protected
	profile.NewHandler(protectedRouter, profileService, ctx)
	filter.NewHandler(protectedRouter, filterService)
	recommendations.NewHandler(protectedRouter, recommendationService)
	like.NewHandler(protectedRouter, likeService)
	match.NewHandler(protectedRouter, matchService)
	chat.NewHandler(protectedRouter, chatService)
	//middlewares
	authMiddleware := middleware.NewAuthMiddleware(*config)
	checkBlockedUserMiddleware := middleware.NewCheckBlockedUserMiddleware(userRepository)
	protectedStackMiddleware := middleware.Chain(
		authMiddleware,
		checkBlockedUserMiddleware,
	)
	globalStackMiddleware := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)
	//routing
	mainRouter.Handle("/dict/", publicRouter)
	mainRouter.Handle("/auth/", publicRouter)
	mainRouter.Handle("/swagger/", publicRouter)
	mainRouter.Handle("/chat/ws", publicRouter)
	mainRouter.Handle("/notifier/ws", publicRouter)
	mainRouter.Handle("/", protectedStackMiddleware(protectedRouter))
	//start-server
	server := http.Server{
		Addr:    ":8081",
		Handler: globalStackMiddleware(mainRouter),
	}
	log.Printf("Server start on %s port. Time: %.3fs\n", server.Addr, time.Since(now).Seconds())
	server.ListenAndServe()

}
