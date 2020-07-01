package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/tludlow/URL-GO/internal/database"
	"github.com/tludlow/URL-GO/internal/router"
)


func main() {
	ctx := context.Background()

	//Load the environment variables from .env into the OS
	godotenv.Load()

	//Connect to the database
	err := database.New(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer database.DB.Close(ctx)
	log.Println("Connected successfully to the database")

	//Create the router
	router := router.New()
	router.Run(":8080")
}