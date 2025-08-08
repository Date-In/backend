package main

import (
	"dating_service/configs"
	"dating_service/internal/auth"
	"dating_service/internal/cache"
	"dating_service/internal/user"
	db2 "dating_service/pkg/db"
	"dating_service/pkg/middleware"
	"log"
	"net/http"
	"time"
)

func main() {
	now := time.Now()
	config := configs.NewConfig()
	db := db2.NewDb(config)
	router := http.NewServeMux()
	refCache, err := cache.NewReferenceCache(db)
	if err != nil {
		panic(err)
	}
	userRepository := user.NewUserRepository(db)
	authService := auth.NewAuthService(config, userRepository, refCache)
	auth.NewAuthHandler(router, authService)

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
