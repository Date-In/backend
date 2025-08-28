package main

import (
	"context"
	"dating_service/configs"
	"dating_service/internal/action"
	"dating_service/internal/activity"
	"dating_service/internal/auth"
	"dating_service/internal/cache"
	"dating_service/internal/chat"
	"dating_service/internal/filestorage"
	"dating_service/internal/filter"
	"dating_service/internal/like"
	"dating_service/internal/match"
	"dating_service/internal/notifier"
	"dating_service/internal/photo"
	"dating_service/internal/profile"
	"dating_service/internal/recommendations"
	"dating_service/internal/user"
	"dating_service/pkg/JWT"
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
	//cache
	refCache, err := cache.NewReferenceCache(db)
	if err != nil {
		panic(err)
	}
	//repository
	userRepository := user.NewUserRepository(db)
	photoRepository := photo.NewPhotoRepository(db)
	filterRepository := filter.NewFilterRepository(db)
	actionsRepository := action.NewActionsRepository(db)
	likeRepository := like.NewLikeRepository(db)
	matchRepository := match.NewMatchRepository(db)
	chatRepository := chat.NewChatRepository(db)
	activityRepository := activity.NewActivityRepository(db)
	//ws-hub
	notifierHub := notifier.NewHub()
	go notifierHub.Run()
	//service
	photoS3Service := filestorage.NewS3FileStorage(photoS3Client, config)
	photoService := photo.NewPhotoService(photoRepository, photoS3Service)
	userService := user.NewUserService(userRepository)
	authService := auth.NewAuthService(userService, refCache, tokenGenerator)
	profileService := profile.NewProfileService(userService, photoService, refCache)
	filterService := filter.NewFilterService(filterRepository)
	recommendationService := recommendations.NewRecommendationService(userService, filterService)
	actionService := action.NewActionsService(userService, actionsRepository)
	matchService := match.NewMatchService(matchRepository)
	likeService := like.NewLikeService(likeRepository, userService, matchService)
	activityService := activity.NewActivityService(activityRepository)
	notifierService := notifier.NewNotifierService(notifierHub, activityService, matchService)
	//background tasks
	go func() {
		ticker := time.NewTicker(24 * time.Hour)
		defer ticker.Stop()
		for {
			actionService.ChangeStatusToNonActive()
			<-ticker.C
		}
	}()

	//handler-public
	fileServer := http.FileServer(http.Dir("./"))
	publicRouter.Handle("/swagger/oas.yaml", http.StripPrefix("/swagger/", fileServer))
	publicRouter.HandleFunc("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/oas.yaml"),
	))
	auth.NewAuthHandler(publicRouter, authService)
	notifier.NewNotifyHandler(publicRouter, notifierService, config)
	chat.NewChatHandlerWs(publicRouter, chatRepository, matchService, config)
	//handler-protected
	profile.NewProfileHandler(protectedRouter, profileService, ctx)
	filter.NewFilterHandler(protectedRouter, filterService)
	recommendations.NewRecommendationHandler(protectedRouter, recommendationService)
	like.NewLikeHandler(protectedRouter, likeService)
	match.NewMatchHandler(protectedRouter, matchService)
	chat.NewChatHandler(protectedRouter, chatRepository, matchService)
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
