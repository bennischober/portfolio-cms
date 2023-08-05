package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"backend/main/internal/api/private"
	"backend/main/internal/config"
	"backend/main/internal/data"
)

func main() {
	// Load environment variables
	err := godotenv.Load(".env.local")
	if err != nil {
		fmt.Printf("Error loading .env file")
	}

	logger := log.New(os.Stdout, "backend", log.LstdFlags|log.Lshortfile)

	// Load config
	config, err := config.LoadConfig("config.yml", logger)
	if err != nil {
		logger.Fatalf("Error loading config: %v", err)
	}

	// Create data store
	store, err := data.NewDataStore(context.Background(), config)
	if err != nil {
		logger.Fatalf("Error creating data store: %v", err)
	}

	// Create gin router
	r := gin.Default()
	

	// Create private api
	privateAPI := private.NewPrivateAPI(store, config)
	privateAPI.SetupRoutes(r)

	// Run server
	logger.Fatal(http.ListenAndServe(":" + config.Server.Port, r))
}
