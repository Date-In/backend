package main

import (
	"dating_service/configs"
	_ "dating_service/docs"
	"dating_service/internal/auth"
	"dating_service/internal/cache"
	"dating_service/internal/filter"
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

// @title API для Сервиса Знакомств
// @version 1.0
// @description Это серверная часть для приложения знакомств. Все эндпоинты, требующие авторизации, ожидают JWT токен в заголовке 'Authorization: Bearer {token}'.
// @termsOfService http://swagger.io/terms/

// @contact.name Ваше Имя
// @contact.email ваш.email@example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8081

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	now := time.Now()
	config := configs.NewConfig()
	tokenGenerator := JWT.NewJWT(config.SecretToken.Token)
	db := db2.NewDb(config)
	router := http.NewServeMux()
	router.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	refCache, err := cache.NewReferenceCache(db)
	if err != nil {
		panic(err)
	}
	//repository
	userRepository := user.NewUserRepository(db)
	photoRepository := photo.NewPhotoRepository(db)
	filterRepository := filter.NewFilterRepository(db)
	//service
	authService := auth.NewAuthService(userRepository, refCache, tokenGenerator)
	profileService := profile.NewProfileService(userRepository, photoRepository, refCache)
	photoService := photo.NewPhotoService(photoRepository)
	filterService := filter.NewFilterService(filterRepository)
	recommendationService := recommendations.NewRecommendationService(userRepository, filterRepository)
	//handler
	profile.NewProfileHandler(router, profileService, config)
	auth.NewAuthHandler(router, authService)
	photo.NewPhotoHandler(router, photoService)
	filter.NewFilterHandler(router, filterService, config)
	recommendations.NewRecommendationHandler(router, recommendationService, config)

	stackMiddleware := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)
	server := http.Server{
		Addr:    ":8081",
		Handler: stackMiddleware(router),
	}
	log.Printf("Server start on %s port. Time: %.3fs\n", server.Addr, time.Since(now).Seconds())
	server.ListenAndServe()

}
