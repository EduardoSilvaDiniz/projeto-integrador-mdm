package entities

import "gorm.io/gorm"

type Present struct {
	gorm.Model
	Meeting    Meeting    `gorm:"no null"`
	Associated Associated `gorm:"no null"`
	Date       string     `gorm:"no null"`
	Present    bool       `gorm:"no null"`
}
