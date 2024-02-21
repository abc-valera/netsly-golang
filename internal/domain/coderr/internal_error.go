package coderr

import (
	"bytes"
	"fmt"
)

type Internal struct {
	Code   Code   // Code provides general information about the error
	Caller string // Caller provides additional context about error's location
	Err    error  // Err is a nested error
}

func NewInternal(err error) error {
	return &Internal{
		Code:   CodeInternal,
		Caller: caller(2),
		Err:    err,
	}
}

func (e *Internal) Error() string {
	var buf bytes.Buffer

	if e.Caller != "" {
		fmt.Fprintf(&buf, "%s ", e.Caller)
	}

	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		if e.Code != "" {
			fmt.Fprintf(&buf, "<%s> ", e.Code)
		}
		buf.WriteString("Internal error")
	}

	return buf.String()
}
