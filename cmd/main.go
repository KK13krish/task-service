package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"task-service/config"
	"task-service/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	config.ConnectDatabase()

	r := routes.SetupRouter()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// test pipeline trigger
