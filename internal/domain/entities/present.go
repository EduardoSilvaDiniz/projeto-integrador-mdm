package entities

import (
	"time"

	"gorm.io/gorm"
)

type Present struct {
	gorm.Model
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Meeting    Meeting        `gorm:"no null"`
	Associated Associated     `gorm:"no null"`
	Date       string         `gorm:"no null"`
	Present    bool           `gorm:"no null"`
}
