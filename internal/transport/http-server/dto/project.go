package dto

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `json:"deletedAt"`
	Name          string         `json:"name"`
	PaymentAmount float32        `json:"paymentAmount"`
}
