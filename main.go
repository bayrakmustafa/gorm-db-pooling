package main

import (
	"gorm-db-pooling/Config"
	"gorm-db-pooling/Models"
	"gorm-db-pooling/Routes"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize Echo instance
	e := echo.New()

	/// Connect to the database
	Config.DatabaseInit()
	defer Config.GetDB().DB()

	// Perform migrations using AutoMigrate
	db := Config.GetDB()
	err := db.AutoMigrate(&Models.Course{})
	if err != nil {
		panic(err)
	}

	// Set up Routes
	Routes.SetupRoutes(e)

	// Start the server
	serverPort := os.Getenv("SERVER_PORT")
	e.Logger.Fatal(e.Start(":" + serverPort))
}
