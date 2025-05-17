package entities

import (
	"time"

	"gorm.io/gorm"
)

type Associated struct {
	gorm.Model
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	CPF           string         `gorm:"not null;unique;primaryKey"`
	Name          string         `gorm:"not null"`
	DateBirth     string         `gorm:"not null"`
	MaritalStatus string         `gorm:"not null"`
}
