package entities

import (
	"reflect"
	"strings"

	"gorm.io/gorm"
)

type MaritalStatus string

const (
	Single   MaritalStatus = "single"
	Married  MaritalStatus = "married"
	Divorced MaritalStatus = "divorced"
)

type Associated struct {
	gorm.Model
	CPF           string        `json:"cpf"           gorm:"no null;unique"`
	Name          string        `json:"name"          gorm:"not null"`
	DateBirth     string        `json:"dateBirth"     gorm:"not null"`
	MaritalStatus MaritalStatus `json:"maritalStatus" gorm:"not null"`
}

func (a Associated) IsValid() []string {
	var error []string
	v := reflect.ValueOf(a)
	t := reflect.TypeOf(a)

	for i := range v.NumField() {
		fieldValue := v.Field(i)
		fieldName := t.Field(i).Name

		if fieldValue.Kind() == reflect.String && strings.TrimSpace(fieldValue.String()) == "" {
			error = append(error, "campo "+fieldName+" esta vazio")
		}

	}
	return error
}
