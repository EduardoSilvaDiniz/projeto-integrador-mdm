package service

import (
	"encoding/json"
	"errors"
	"io"
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

		if fieldValue.Kind() == reflect.String && strings.TrimSpace(fieldValue.String()) == "" || strings.TrimSpace(fieldValue.String()) == "0" {
			listErrors = append(listErrors, "campo "+fieldName+" esta vazio")
		}
	}
	if listErrors != nil {
		return errors.New(strings.Join(listErrors, " "))
	}

	return nil
}

func serialization[T any, D any](body io.ReadCloser) (D, error) {
	var dto T
	var result D

	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return result, err
	}

	if err := IsValid(dto); err != nil {
		return result, err
	}

	dstPtr := reflect.New(reflect.TypeOf((*D)(nil)).Elem())
	dstVal := dstPtr.Elem()
	srcVal := reflect.ValueOf(dto)

	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}

	for i := range dstVal.NumField() {
		field := dstVal.Type().Field(i).Name
		srcField := srcVal.FieldByName(field)
		if srcField.IsValid() && srcField.Type().AssignableTo(dstVal.Field(i).Type()) {
			dstVal.Field(i).Set(srcField)
		}
	}

	result = dstPtr.Interface().(D)
	return result, nil
}
