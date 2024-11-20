package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/justindoan/photobooth-backend/controllers"
	"github.com/justindoan/photobooth-backend/dto"
)

// ImageHandler handles HTTP requests related to images
type ImageHandler struct {
	imageController *controllers.ImageController
	routes          []Route
}

// NewImageHandler creates a new instance of ImageHandler
func NewImageHandler(imageController *controllers.ImageController) *ImageHandler {
	h := &ImageHandler{imageController: imageController}
	h.routes = []Route{
		{http.MethodPost, "process", h.HandleProcessImage},
		{http.MethodGet, "list", h.HandleListImages},
		{http.MethodGet, "progress", h.HandleProgress},
	}
	return h
}

// ServeHTTP implements the http.Handler interface
func (h *ImageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Loop through defined routes and call the appropriate handler
	for _, route := range h.routes {
		if r.Method == route.method && r.URL.Path == route.path {
			route.handler(w, r)
			return
		}
	}
	http.NotFound(w, r)
}

// HandleProcessImage handles the HTTP request for processing an image
func (h *ImageHandler) HandleProcessImage(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body into a ProcessImageRequest struct
	var processImageRequest dto.ProcessImageRequest
	err := json.NewDecoder(r.Body).Decode(&processImageRequest)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the request
	if processImageRequest.Data == "" {
		http.Error(w, "Image data is required", http.StatusBadRequest)
		return
	}

	// Call the controller to process the image
	result := h.imageController.ProcessImage(r.Context(), processImageRequest)

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Image processed", "result": result})
}

// HandleListImages handles the HTTP request for listing all images
func (h *ImageHandler) HandleListImages(w http.ResponseWriter, r *http.Request) {
	// Retrieve the list of images from the controller
	images, err := h.imageController.ListImages(r.Context())
	if err != nil {
		http.Error(w, "Failed to list images", http.StatusInternalServerError)
		return
	}

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(images)
}

func (h *ImageHandler) HandleProgress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp, err := http.Get("http://127.0.0.1:7860/sdapi/v1/progress")
	if err != nil {
		http.Error(w, "Failed to get progress", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(body); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}
