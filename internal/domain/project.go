package domain

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Id            int     `json:"id"`
	Name          string  `json:"name"`
	PaymentAmount float32 `json:"paymentAmount"`
}
