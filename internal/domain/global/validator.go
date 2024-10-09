package global

import (
	"errors"
	"fmt"

	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"github.com/abc-valera/netsly-golang/internal/domain/util/enum"
	"github.com/go-playground/validator/v10"
)

// IValidator is a wrapper for go-playground/validator/v10 and is a domain dependency.
// Should be created as a singleton and used across the entity layer.
type IValidator interface {
	Struct(s any) error
	Var(field any, tag string) error
}

type validate struct {
	validate *validator.Validate
}

func newValidator() IValidator {
	v := validator.New(validator.WithRequiredStructEnabled())

	// Define and register custom validation functions here:

	// IEnum validation
	validateIEnum := func(fl validator.FieldLevel) bool {
		value, ok := fl.Field().Interface().(enum.IEnum)
		if !ok {
			Log().Error("IEnum validation failed: field is not an IEnum")
			return false
		}

		return value.IsValid()
	}

	v.RegisterValidation("enum", validateIEnum)

	return validate{
		validate: v,
	}
}

func (v validate) Struct(s any) error {
	if err := v.validate.Struct(s); err != nil {
		validationErrs, ok := err.(validator.ValidationErrors)
		if !ok {
			return coderr.NewInternalErr(errors.New("failed to cast validation error"))
		}

		var returnErr error
		for _, e := range validationErrs {
			if e.Param() != "" {
				returnErr = errors.Join(
					v.createFormattedErrorWithParam(e.Field(), e.Tag(), e.Param()),
					returnErr,
				)
				break // Note: not sure if this is the best way to handle multiple errors
			}

			returnErr = errors.Join(
				v.createFormattedError(e.Field(), e.Tag()),
				returnErr,
			)
			break // Note: not sure if this is the best way to handle multiple errors
		}
		return returnErr
	}
	return nil
}

func (v validate) Var(field any, tag string) error {
	if err := v.validate.Var(field, tag); err != nil {
		if e, ok := err.(validator.FieldError); ok {
			return v.createFormattedError(e.Field(), e.Tag())
		}
	}
	return nil
}

func (validate) createFormattedError(field, tag string) error {
	return coderr.NewCodeMessage(
		coderr.CodeInvalidArgument,
		fmt.Sprintf("%s '%s validation rule' violated", field, tag),
	)
}

func (validate) createFormattedErrorWithParam(field, tag, param string) error {
	return coderr.NewCodeMessage(
		coderr.CodeInvalidArgument,
		fmt.Sprintf("%s '%s %s validation rule' violated", field, tag, param),
	)
}
