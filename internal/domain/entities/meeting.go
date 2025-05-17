package entities

import (
	"time"

	"gorm.io/gorm"
)

type Meeting struct {
	gorm.Model
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	ProjectId int            `gorm:"no null"`
	GroupId   int            `gorm:"no null"`
	HubId     int            `gorm:"no null"`
	Date      string         `gorm:"no null"`
	Time      string         `gorm:"no null"`
}
