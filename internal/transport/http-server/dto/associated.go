package dto

import (
	"chamada-pagamento-system/internal/domain/entities"
	"errors"
	"reflect"
	"strings"
	"time"

	"gorm.io/gorm"
)

type MaritalStatus string

const (
	Single   MaritalStatus = "single"
	Married  MaritalStatus = "married"
	Divorced MaritalStatus = "divorced"
)

type Associated struct {
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
	CPF           string         `json:"cpf,omitempty"`
	Name          string         `json:"name,omitempty"`
	DateBirth     string         `json:"date_birth,omitempty"`
	MaritalStatus MaritalStatus  `json:"marital_status,omitempty"`
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

func (a *Associated) ToEntity() *entities.Associated {
	return &entities.Associated{
		CreatedAt:     a.CreatedAt,
		UpdatedAt:     a.UpdatedAt,
		DeletedAt:     a.DeletedAt,
		CPF:           a.CPF,
		Name:          a.Name,
		DateBirth:     a.DateBirth,
		MaritalStatus: string(a.MaritalStatus),
	}
}
