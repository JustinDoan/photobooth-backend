package repository

import (
	"context"

	"github.com/justindoan/photobooth-backend/models"
	"gorm.io/gorm"
)

// ImageRepository handles database operations for images
type ImageRepository struct {
	db *gorm.DB
}

// NewImageRepository creates a new instance of ImageRepository
func NewImageRepository(db *gorm.DB) *ImageRepository {
	return &ImageRepository{db: db}
}

// Create inserts a new image record into the database
func (r *ImageRepository) Create(ctx context.Context, image *models.Image) error {
	return r.db.WithContext(ctx).Create(image).Error
}

// List retrieves all image records from the database
func (r *ImageRepository) List(ctx context.Context) ([]models.Image, error) {
	var images []models.Image
	err := r.db.WithContext(ctx).Find(&images).Error
	return images, err
}
