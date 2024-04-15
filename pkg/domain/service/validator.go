package service

import (
	"errors"
	"fmt"

	"github.com/abc-valera/netsly-api-golang/pkg/core/coderr"
	"github.com/go-playground/validator/v10"
)

// IValidator is a wrapper for go-playground/validator/v10 and is a domain dependency.
// Should be created as a singleton and used across the entity layer.
type IValidator interface {
	Struct(s interface{}) error
	Var(field interface{}, tag string) error
}

type validate struct {
	validate *validator.Validate
}

func NewValidator() IValidator {
	return validate{
		validate: validator.New(),
	}
}

func (v validate) Struct(s interface{}) error {
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

func (v validate) Var(field interface{}, tag string) error {
	if err := v.validate.Var(field, tag); err != nil {
		if e, ok := err.(validator.FieldError); ok {
			return v.createFormattedError(e.Field(), e.Tag())
		}
	}
	return nil
}

func (v validate) createFormattedError(field, tag string) error {
	return coderr.NewCodeMessage(
		coderr.CodeInvalidArgument,
		fmt.Sprintf("%s '%s validation rule' violated", field, tag),
	)
}

func (v validate) createFormattedErrorWithParam(field, tag, param string) error {
	return coderr.NewCodeMessage(
		coderr.CodeInvalidArgument,
		fmt.Sprintf("%s '%s %s validation rule' violated", field, tag, param),
	)
}
