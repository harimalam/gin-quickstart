package main

import (
	"gin-quickstart/internal/albums"
	"gin-quickstart/internal/auth"
	"gin-quickstart/internal/config"
	"gin-quickstart/internal/db"
	"gin-quickstart/internal/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	Cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	// Database setup
	database, err := db.InitDB(Cfg)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Auth setup
	authRepo := auth.NewRepository(database)
	authService := auth.NewService(authRepo, Cfg)
	authHandler := auth.NewHandler(authService)

	// Albums setup
	albumRepo := albums.NewRepository(database)
	albumService := albums.NewService(albumRepo)
	albumHandler := albums.NewHandler(albumService)

	// Create router and register feature routes.
	router := gin.Default()

	// API v1 group
	apiGroup := router.Group("/api/v1")

	// PUBLIC ROUTES
	publicGroup := apiGroup.Group("/")
	authHandler.RegisterRoutes(publicGroup)

	// PROTECTED ROUTES
	protectedGroup := apiGroup.Group("/")
	protectedGroup.Use(middleware.AuthMiddleware([]byte(Cfg.App.JWTSecret)))
	{
		albumHandler.RegisterRoutes(protectedGroup)
	}

	// 4. Start HTTP server.
	port := ":" + Cfg.App.Port
	router.Run(port)
}
