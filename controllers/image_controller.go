package controllers

import (
	"context"

	"github.com/justindoan/photobooth-backend/dto"
	"github.com/justindoan/photobooth-backend/models"
	"github.com/justindoan/photobooth-backend/repository"
	"github.com/justindoan/photobooth-backend/services"
)

type ImageController struct {
	services     *services.Services
	repositories *repository.Repositories
}

func NewImageController(s *services.Services, r *repository.Repositories) *ImageController {
	return &ImageController{
		services:     s,
		repositories: r,
	}
}

func (c *ImageController) ProcessImage(ctx context.Context, image dto.ProcessImageRequest) string {
	// Responsible for actually processing the image, should return the raw data to the handler for it to send to the client

	// An example of how to use the database here to store some image metadata
	c.repositories.ImageRepository.Create(ctx, &models.Image{
		Name: image.Name,
	})

	// Here we would call the image processing service
	// c.services.ImageService.ProcessImage(ctx, image) for example
	return "Processed Image"
}

func (c *ImageController) ListImages(ctx context.Context) ([]models.Image, error) {
	return c.repositories.ImageRepository.List(ctx)
}
