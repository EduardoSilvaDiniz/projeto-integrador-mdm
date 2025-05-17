package entities

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	Associated     int    `gorm:"no null"`
	Date           string `gorm:"no null"`
	MonthReference string `gorm:"no null"`
}
