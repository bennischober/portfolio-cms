package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"

	"backend/main/internal/api"
	"backend/main/internal/auth"
	"backend/main/internal/config"
	"backend/main/internal/data"
	"backend/main/internal/models"
)

func InitializeLogger() *log.Logger {
	return log.New(os.Stdout, "backend", log.LstdFlags|log.Lshortfile)
}

func LoadConfig(logger *log.Logger) *config.Config {
	err := godotenv.Load(".env.local")
	if err != nil {
		logger.Fatalf("Error loading .env file")
	}

	config, err := config.LoadConfig("config.yml", logger)
	if err != nil {
		logger.Fatalf("Error loading config: %v", err)
	}

	return config
}

func InitializeDataStore(config *config.Config, logger *log.Logger) models.DataStore {
	store, err := data.NewDataStore(context.Background(), config)
	if err != nil {
		logger.Fatalf("Error creating data store: %v", err)
	}
	return store
}

func SetupRoutes(config *config.Config, store models.DataStore, logger *log.Logger) *gin.Engine {
	// Create gin router
	r := gin.Default()
	
	// Create auth service
	tokenService := auth.NewTokenService()
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Authentication.RedisAddress,
		Password: config.Authentication.RedisPassword,
		DB:       config.Authentication.RedisDB,
	})
	redisService := auth.NewRedisService(context.Background(), redisClient)
	authService := auth.NewAuthService(config, store, tokenService, redisService)
	middleware := auth.NewAuthMiddleware(tokenService)

	// Create private and public API
	api := api.NewAPIBuilder(store, config, authService, middleware, tokenService)
	api.SetupRoutes(r)

	return r
}
