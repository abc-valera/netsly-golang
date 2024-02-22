package validation

import (
	"errors"
	"fmt"

	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/go-playground/validator/v10"
)

// Validator is a wrapper for go-playground/validator/v10.
// Should be created as a singleton and used across the application.
type Validator struct {
	validate *validator.Validate
}

func NewValidator() Validator {
	return Validator{
		validate: validator.New(),
	}
}

func (v Validator) Struct(s interface{}) error {
	if err := v.validate.Struct(s); err != nil {
		var returnErr error
		for _, e := range err.(validator.ValidationErrors) {
			if e.Param() != "" {
				returnErr = errors.Join(
					v.createFormattedErrorWithParam(e.Field(), e.Tag(), e.Param()),
					returnErr,
				)
				break // Note: not sure if this is the best way to handle multiple errors
			} else {
				returnErr = errors.Join(
					v.createFormattedError(e.Field(), e.Tag()),
					returnErr,
				)
				break // Note: not sure if this is the best way to handle multiple errors
			}
		}
		return returnErr
	}
	return nil
}

func (v Validator) Var(field interface{}, tag string) error {
	if err := v.validate.Var(field, tag); err != nil {
		if e, ok := err.(validator.FieldError); ok {
			return v.createFormattedError(e.Field(), e.Tag())
		}
	}
	return nil
}

func (v Validator) createFormattedError(field, tag string) error {
	return coderr.NewMessage(
		coderr.CodeInvalidArgument,
		fmt.Sprintf("%s '%s validation rule' violated", field, tag),
	)
}

func (v Validator) createFormattedErrorWithParam(field, tag, param string) error {
	return coderr.NewMessage(
		coderr.CodeInvalidArgument,
		fmt.Sprintf("%s '%s %s validation rule' violated", field, tag, param),
	)
}
