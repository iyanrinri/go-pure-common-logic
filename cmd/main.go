package main

import (
	"go-common-logic/internal/routes"
	"log"
)

func main() {
	router := routes.SetupRoutes()

	log.Println("Starting server on port 8081...")
	if err := router.Run(":8081"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
