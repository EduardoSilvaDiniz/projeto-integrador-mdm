package service

import (
	"errors"
	"reflect"
	"strings"
)

func IsValid(a any) error {
	var listErrors []string
	v := reflect.ValueOf(a)
	t := reflect.TypeOf(a)

	for i := range v.NumField() {
		fieldValue := v.Field(i)
		fieldName := t.Field(i).Name

		if fieldValue.Kind() == reflect.String && strings.TrimSpace(fieldValue.String()) == "" || strings.TrimSpace(fieldValue.String()) == "0"{
			listErrors = append(listErrors, "campo "+fieldName+" esta vazio")
		}
	}
	if listErrors != nil {
		return errors.New(strings.Join(listErrors, " "))
	}

	return nil
}
