package models

import "time"

type Image struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
