package dto

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `json:"deletedAt"`
	Associated     int            `json:"associated"`
	Date           string         `json:"date"`
	MonthReference string         `json:"monthReference"`
}
