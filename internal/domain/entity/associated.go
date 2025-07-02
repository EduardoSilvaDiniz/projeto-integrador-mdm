package entity

import (
	"errors"
	"reflect"
	"strings"

	"github.com/google/uuid"
)

type MaritalStatus string

const (
	Single   MaritalStatus = "single"
	Married  MaritalStatus = "married"
	Divorced MaritalStatus = "divorced"
)

type Associated struct {
	ID            uuid.UUID     `json:"id"`
	CPF           string        `json:"cpf,omitempty"`
	Name          string        `json:"name,omitempty"`
	DateBirth     string        `json:"date_birth,omitempty"`
	MaritalStatus MaritalStatus `json:"marital_status,omitempty"`
}

func (a Associated) IsValid() error {
	var listErrors []string
	v := reflect.ValueOf(a)
	t := reflect.TypeOf(a)

	for i := range v.NumField() {
		fieldValue := v.Field(i)
		fieldName := t.Field(i).Name

		if fieldValue.Kind() == reflect.String && strings.TrimSpace(fieldValue.String()) == "" {
			listErrors = append(listErrors, "campo "+fieldName+" esta vazio")
		}
	}
	if listErrors != nil {
		return errors.New(strings.Join(listErrors, " "))
	}

	return nil
}
