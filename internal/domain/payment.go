package domain

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	Id             int    `json:"id"`
	Associated     int    `json:"associated"`
	Date           string `json:"date"`
	MonthReference string `json:"monthReference"`
}
