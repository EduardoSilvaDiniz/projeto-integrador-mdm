package service

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func IsValid(a any) error {
	var listErrors []string
	v := reflect.ValueOf(a)
	t := reflect.TypeOf(a)

	for i := range v.NumField() {
		fieldValue := v.Field(i)
		fieldName := t.Field(i).Name

		if fieldValue.Kind() == reflect.String && strings.TrimSpace(fieldValue.String()) == "" ||
			strings.TrimSpace(fieldValue.String()) == "0" {
			listErrors = append(listErrors, "campo "+fieldName+" esta vazio")
		}
	}
	if listErrors != nil {
		return errors.New(strings.Join(listErrors, " "))
	}

	return nil
}

var validate = validator.New()

func ValidateStruct[T any](input T) error {
	if err := validate.Struct(input); err != nil {
		var listErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			// fmt.Printf("Erro no campo '%s': %s\n", err.Field(), err.ActualTag())
			listErrors = append(listErrors, "campo invalido", err.Field(), err.ActualTag())
		}

		return errors.New(strings.Join(listErrors, " "))
	}

	return nil
}

func FormatValidationError(err error) []string {
	var errs []string
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, ve := range validationErrors {
			errs = append(
				errs,
				fmt.Sprintf("Campo '%s' inválido: regra '%s'", ve.Field(), ve.Tag()),
			)
		}
	} else {
		errs = append(errs, err.Error())
	}
	return errs
}

func decodeJson(body io.ReadCloser) ([]byte, error) {
	data, err := io.ReadAll(body)
	if err != nil {
		slog.Error("Falha ao decodificar corpo da requisição JSON")
		return nil, err
	}

	slog.Debug("corpo recebido:", "data", string(data))
	return data, nil
}
