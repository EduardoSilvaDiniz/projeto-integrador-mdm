package domain

import "gorm.io/gorm"

type MaritalStatus string

const (
	Single   MaritalStatus = "single"
	Married  MaritalStatus = "married"
	Divorced MaritalStatus = "divorced"
)

type Associated struct {
	gorm.Model
	Cpf           string        `json:"cpf"`
	Name          string        `json:"name"`
	DateBirth     string        `json:"dateBirth"`
	MaritalStatus MaritalStatus `json:"maritalStatus"`
}
