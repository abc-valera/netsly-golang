package codeerr

import (
	"bytes"
	"fmt"
)

type InternalErr struct {
	Code Code   // Code provides general information about the error
	Op   string // Op (operation) provides additional context about error's location
	Err  error  // Err is a nested error
}

func NewInternal(op string, err error) error {
	return &InternalErr{
		Code: CodeInternal,
		Op:   op,
		Err:  err,
	}
}

// Error returns the string representation of the Error
func (e *InternalErr) Error() string {
	var buf bytes.Buffer

	if e.Op != "" {
		fmt.Fprintf(&buf, "%s: ", e.Op)
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
