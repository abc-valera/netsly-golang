package coderr

import (
	"reflect"
)

func CheckIfStructHasEmptyFields(data any) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = NewInternalString("Panic occured: " + r.(string))
		}
	}()

	if data == nil {
		return NewInternalString("Data is nil")
	}

	structVal := reflect.ValueOf(data)
	if structVal.Kind() != reflect.Struct {
		return NewInternalString("Data is not a struct")
	}

	for i := range structVal.NumField() {
		fieldVal := reflect.ValueOf(structVal.Field(i).Interface())

		if !fieldVal.IsValid() {
			return NewInternalString("Field " + structVal.Type().Field(i).Name + " is zero value")
		}

		switch fieldVal.Kind() {
		case reflect.Pointer, reflect.Interface:
			if fieldVal.IsNil() {
				return NewInternalString("Field " + structVal.Type().Field(i).Name + " is nil")
			}
		default:
			continue
		}
	}

	return nil
}
