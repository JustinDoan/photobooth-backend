package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/justindoan/photobooth-backend/controllers"
	"github.com/justindoan/photobooth-backend/dto"
)

type ImageHandler struct {
	imageController *controllers.ImageController
	routes          []Route
}

func NewImageHandler(imageController *controllers.ImageController) *ImageHandler {
	h := &ImageHandler{imageController: imageController}
	h.routes = []Route{
		{http.MethodPost, "process", h.HandleProcessImage},
		{http.MethodGet, "list", h.HandleListImages},
	}
	return h
}

func (h *ImageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range h.routes {
		if r.Method == route.method && r.URL.Path == route.path {
			route.handler(w, r)
			return
		}
	}
	http.NotFound(w, r)
}

func (h *ImageHandler) HandleProcessImage(w http.ResponseWriter, r *http.Request) {
	// Responsible for parsing the request body, and calling the appropriate controller method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var processImageRequest dto.ProcessImageRequest

	err := json.NewDecoder(r.Body).Decode(&processImageRequest)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if processImageRequest.Data == "" {
		http.Error(w, "Image data is required", http.StatusBadRequest)
		return
	}

	// Pass the context and image to the image controller
	result := h.imageController.ProcessImage(r.Context(), processImageRequest)

	// For now, just send a simple response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Image processed", "result": result})
}

func (h *ImageHandler) HandleListImages(w http.ResponseWriter, r *http.Request) {

	images, err := h.imageController.ListImages(r.Context())
	if err != nil {
		http.Error(w, "Failed to list images", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(images)
}
