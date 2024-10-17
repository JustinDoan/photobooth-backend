package repository

import (
	"gorm.io/gorm"
)

type Repositories struct {
	ImageRepository *ImageRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		ImageRepository: NewImageRepository(db),
	}
}
