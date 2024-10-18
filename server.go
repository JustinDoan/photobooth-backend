package main

import (
	"net/http"
	"strings"

	"github.com/justindoan/photobooth-backend/controllers"
	"github.com/justindoan/photobooth-backend/handlers"
	"github.com/justindoan/photobooth-backend/repository"
	"github.com/justindoan/photobooth-backend/services"
)

// Server represents the HTTP server and its handlers
type Server struct {
	handlers map[string]http.Handler
}

// NewServer creates a new instance of Server with configured handlers
func NewServer(services *services.Services, repositories *repository.Repositories) *Server {
	// Create controllers and handlers
	imageController := controllers.NewImageController(services, repositories)
	imageHandler := handlers.NewImageHandler(imageController)

	// Configure the server with handlers
	return &Server{
		handlers: map[string]http.Handler{
			"/api/images/": imageHandler,
		},
	}
}

// SetupRoutes configures the routing for the server
func (s *Server) SetupRoutes() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Iterate through registered handlers and match the URL path
		for prefix, handler := range s.handlers {
			if strings.HasPrefix(r.URL.Path, prefix) {
				// Strip the prefix from the path and pass the request to the handler
				http.StripPrefix(prefix, handler).ServeHTTP(w, r)
				return
			}
		}
		// If no handler is found, return 404 Not Found
		http.NotFound(w, r)
	})
}
