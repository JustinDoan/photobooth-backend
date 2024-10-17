package database

import (
	"github.com/justindoan/photobooth-backend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto Migrate the schema
	err = db.AutoMigrate(&models.Image{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
