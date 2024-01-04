package codeerr

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
)

type Internal struct {
	Code   Code   // Code provides general information about the error
	Caller string // Caller provides additional context about error's location
	Err    error  // Err is a nested error
}

func NewInternal(err error) error {
	// Get the caller program counter and line number
	pc, _, line, _ := runtime.Caller(1)
	// Get the function name
	split := strings.Split(runtime.FuncForPC(pc).Name(), "/")
	funcName := strings.Join(split[4:], "/")
	return &Internal{
		Code:   CodeInternal,
		Caller: fmt.Sprintf("%s:%d", funcName, line),
		Err:    err,
	}
}

func (e *Internal) Error() string {
	var buf bytes.Buffer

	if e.Caller != "" {
		fmt.Fprintf(&buf, "[%s] ", e.Caller)
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
