// Package controllers contains the business logic for handling requests
package controllers

import (
	"context"

	"github.com/justindoan/photobooth-backend/dto"
	"github.com/justindoan/photobooth-backend/models"
	"github.com/justindoan/photobooth-backend/repository"
	"github.com/justindoan/photobooth-backend/services"
)

// ImageController handles the business logic for image-related operations
type ImageController struct {
	services     *services.Services
	repositories *repository.Repositories
}

// NewImageController creates a new instance of ImageController
func NewImageController(s *services.Services, r *repository.Repositories) *ImageController {
	return &ImageController{
		services:     s,
		repositories: r,
	}
}

// ProcessImage handles the logic for processing an image
func (c *ImageController) ProcessImage(ctx context.Context, image dto.ProcessImageRequest) string {
	// TODO: Implement actual image processing logic

	// Store image metadata in the database
	c.repositories.ImageRepository.Create(ctx, &models.Image{
		Name: image.Name,
	})

	// Here we would call the image processing service
	// c.services.ImageService.ProcessImage(ctx, image) for example
	return "Processed Image"
}

// ListImages retrieves a list of all images from the database
func (c *ImageController) ListImages(ctx context.Context) ([]models.Image, error) {
	return c.repositories.ImageRepository.List(ctx)
}
