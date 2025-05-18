package dto

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
	Name          string         `json:"name,omitempty"`
	PaymentAmount float32        `json:"payment_amount,omitempty"`
}
