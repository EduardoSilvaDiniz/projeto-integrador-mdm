package entities

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	Associated     int            `gorm:"no null"`
	Date           string         `gorm:"no null"`
	MonthReference string         `gorm:"no null"`
}
