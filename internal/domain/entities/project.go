package entities

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name          string  `gorm:"no null"`
	PaymentAmount float32 `gorm:"no null"`
}
