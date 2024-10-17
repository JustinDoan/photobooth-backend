package main

import (
	"net/http"
	"strings"

	"github.com/justindoan/photobooth-backend/controllers"
	"github.com/justindoan/photobooth-backend/handlers"
	"github.com/justindoan/photobooth-backend/repository"
	"github.com/justindoan/photobooth-backend/services"
)

type Server struct {
	handlers map[string]http.Handler
}

func NewServer(services *services.Services, repositories *repository.Repositories) *Server {
	imageController := controllers.NewImageController(services, repositories)
	imageHandler := handlers.NewImageHandler(imageController)

	return &Server{
		handlers: map[string]http.Handler{
			"/api/images/": imageHandler,
		},
	}
}

func (s *Server) SetupRoutes() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for prefix, handler := range s.handlers {
			if strings.HasPrefix(r.URL.Path, prefix) {
				http.StripPrefix(prefix, handler).ServeHTTP(w, r)
				return
			}
		}
		http.NotFound(w, r)
	})
}
