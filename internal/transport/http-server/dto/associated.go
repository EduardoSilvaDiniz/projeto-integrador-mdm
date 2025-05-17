package dto

import (
	"reflect"
	"strings"
)

type MaritalStatus string

const (
	Single   MaritalStatus = "single"
	Married  MaritalStatus = "married"
	Divorced MaritalStatus = "divorced"
)

type Associated struct {
	CPF           string        `json:"cpf"`
	Name          string        `json:"name"`
	DateBirth     string        `json:"dateBirth"`
	MaritalStatus MaritalStatus `json:"maritalStatus"`
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
