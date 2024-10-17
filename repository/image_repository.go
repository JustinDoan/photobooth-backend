package repository

import (
	"context"

	"github.com/justindoan/photobooth-backend/models"
	"gorm.io/gorm"
)

type ImageRepository struct {
	db *gorm.DB
}

func NewImageRepository(db *gorm.DB) *ImageRepository {
	return &ImageRepository{db: db}
}

func (r *ImageRepository) Create(ctx context.Context, image *models.Image) error {
	return r.db.WithContext(ctx).Create(image).Error
}

func (r *ImageRepository) List(ctx context.Context) ([]models.Image, error) {
	var images []models.Image
	err := r.db.WithContext(ctx).Find(&images).Error
	return images, err
}
