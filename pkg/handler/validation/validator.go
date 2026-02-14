package validation

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	once sync.Once
	v    *validator.Validate
)

func get() *validator.Validate {
	once.Do(func() {
		v = validator.New()
	})
	return v
}

func Validate[T any](payload T) error {
	return get().Struct(payload)
}
