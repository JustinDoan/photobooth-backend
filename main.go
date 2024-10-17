package main

import (
	"log"
	"net/http"

	"github.com/justindoan/photobooth-backend/database"
	"github.com/justindoan/photobooth-backend/repository"
	"github.com/justindoan/photobooth-backend/services"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// Create Services and Repositories
	services := services.NewServices()
	repositories := repository.NewRepositories(db)

	// Create and setup the server
	server := NewServer(services, repositories)
	handler := server.SetupRoutes()

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
