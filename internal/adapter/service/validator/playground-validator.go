package validator

import (
	"errors"
	"fmt"

	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
	"github.com/go-playground/validator/v10"
)

type playgroundValidator struct {
	validate *validator.Validate
}

func NewPlaygroundValidator() service.IValidator {
	return playgroundValidator{
		validate: validator.New(),
	}
}

func (v playgroundValidator) Struct(s interface{}) error {
	if err := v.validate.Struct(s); err != nil {
		var returnErr error
		for _, e := range err.(validator.ValidationErrors) {
			field := e.Field()
			tag := e.Tag()
			param := e.Param()
			if param != "" {
				returnErr = errors.Join(
					coderr.NewMessage(
						coderr.CodeInvalidArgument,
						fmt.Sprintf("%s '%s %s validation rule' violated", field, tag, param),
					),
					returnErr,
				)
			} else {
				returnErr = errors.Join(
					coderr.NewMessage(
						coderr.CodeInvalidArgument,
						fmt.Sprintf("%s '%s validation rule' violated", field, tag),
					),
					returnErr,
				)
			}
		}
		return returnErr
	}
	return nil
}

func (v playgroundValidator) Var(field interface{}, tag string) error {
	if err := v.validate.Var(field, tag); err != nil {
		if e, ok := err.(validator.FieldError); ok {
			return coderr.NewMessage(
				coderr.CodeInvalidArgument,
				fmt.Sprintf("%s '%s validation rule' violated", e.Field(), e.Tag()),
			)
		}
	}
	return nil
}
