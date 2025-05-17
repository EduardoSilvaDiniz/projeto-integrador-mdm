package entities

import "gorm.io/gorm"

type Associated struct {
	gorm.Model
	CPF           string `gorm:"no null;unique"`
	Name          string `gorm:"not null"`
	DateBirth     string `gorm:"not null"`
	MaritalStatus string `gorm:"not null"`
}
