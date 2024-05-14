package validator

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/core/optional"
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
	val := validator.New(validator.WithRequiredStructEnabled())

	optionalCustomTypeFunc := func(field reflect.Value) interface{} {
		opt := field.Interface()
		switch t := opt.(type) {
		case optional.Optional[string]:
			if t.IsPresent() {
				return t.Value()
			}
		case optional.Optional[int]:
			if t.IsPresent() {
				return t.Value()
			}
		case optional.Optional[int32]:
			if t.IsPresent() {
				return t.Value()
			}
		case optional.Optional[int64]:
			if t.IsPresent() {
				return t.Value()
			}
		case optional.Optional[uint]:
			if t.IsPresent() {
				return t.Value()
			}
		case optional.Optional[uint32]:
			if t.IsPresent() {
				return t.Value()
			}
		case optional.Optional[uint64]:
			if t.IsPresent() {
				return t.Value()
			}
		case optional.Optional[bool]:
			if t.IsPresent() {
				return t.Value()
			}
		case optional.Optional[float32]:
			if t.IsPresent() {
				return t.Value()
			}
		case optional.Optional[float64]:
			if t.IsPresent() {
				return t.Value()
			}
		case optional.Optional[byte]:
			if t.IsPresent() {
				return t.Value()
			}
		default:
			return nil
		}
		return nil
	}

	val.RegisterCustomTypeFunc(
		optionalCustomTypeFunc,
		optional.Optional[string]{},
		optional.Optional[int]{},
		optional.Optional[int32]{},
		optional.Optional[int64]{},
		optional.Optional[uint]{},
		optional.Optional[uint32]{},
		optional.Optional[uint64]{},
		optional.Optional[bool]{},
		optional.Optional[float32]{},
		optional.Optional[float64]{},
		optional.Optional[byte]{},
	)

	return validate{
		validate: val,
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
