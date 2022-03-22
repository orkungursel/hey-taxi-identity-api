package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

func Validate(r interface{}) error {
	v := validator.New()
	if err := v.Struct(r); err != nil {
		return errors.Wrap(err, "validation error")
	}

	return nil
}
