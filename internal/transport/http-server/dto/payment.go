package dto

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
	Associated     int            `json:"associated,omitempty"`
	Date           string         `json:"date,omitempty"`
	MonthReference string         `json:"month_reference,omitempty"`
}
