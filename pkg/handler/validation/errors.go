package validation

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ValidationErrors map[string]string

func ToValidationErrors(err error) ValidationErrors {
	var ve validator.ValidationErrors
	if !errors.As(err, &ve) {
		return nil
	}

	fields := make(map[string]string)
	for _, fe := range ve {
		fields[fe.Field()] = fe.Tag()
	}
	return fields
}
