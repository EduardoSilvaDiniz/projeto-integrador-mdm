package entities

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Name          string         `gorm:"no null"`
	PaymentAmount float32        `gorm:"no null"`
}
