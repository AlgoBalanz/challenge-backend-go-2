package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Category  string    `json:"category"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}
