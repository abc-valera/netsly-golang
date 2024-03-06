package coderr

import (
	"bytes"
	"fmt"
)

type codeMessage struct {
	Code    Code   // Code provides general information about the error
	Message string // Message provides additional context in human-readable form
}

func NewCodeMessage(code Code, message string) error {
	return &codeMessage{
		Code:    code,
		Message: message,
	}
}

func NewCodeError(code Code, message error) error {
	return &codeMessage{
		Code:    code,
		Message: message.Error(),
	}
}

func NewCodeMessageError(code Code, message string, err error) error {
	return &codeMessage{
		Code:    code,
		Message: fmt.Sprintf("%s: %s", message, err.Error()),
	}
}

// Error returns the string representation of the Error
func (e codeMessage) Error() string {
	var buf bytes.Buffer

	if e.Code != "" {
		fmt.Fprintf(&buf, "<%s> ", e.Code)
	}
	buf.WriteString(e.Message)

	return buf.String()
}
