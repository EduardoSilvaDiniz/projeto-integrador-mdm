package entities

import "gorm.io/gorm"

type Meeting struct {
	gorm.Model
	ProjectId int    `gorm:"no null"`
	GroupId   int    `gorm:"no null"`
	HubId     int    `gorm:"no null"`
	Date      string `gorm:"no null"`
	Time      string `gorm:"no null"`
}
