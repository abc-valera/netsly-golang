package codeerr

import (
	"bytes"
	"errors"
	"fmt"
)

type MessageErr struct {
	Code    Code   // Code provides general information about the error
	Message string // Message provides additional context in human-readable form
	Err     error  // Err is a nested error
}

func NewMessageErr(code Code, message string) error {
	return &MessageErr{
		Code:    code,
		Message: message,
	}
}

// Error returns the string representation of the Error
func (e *MessageErr) Error() string {
	var buf bytes.Buffer

	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		if e.Code != "" {
			fmt.Fprintf(&buf, "<%s> ", e.Code)
		}
		buf.WriteString(e.Message)
	}

	return buf.String()
}

// ErrorMessage returns the message of the error, if available. Otherwise returns a error message.
// If error is nil returns an empty string.
func ErrorMessage(err error) string {
	var messageErrTarget *MessageErr
	if err == nil {
		return ""
	} else if e, ok := err.(*MessageErr); ok && e.Message != "" {
		return e.Message
	} else if errors.As(err, &messageErrTarget) {
		return messageErrTarget.Message
	} else if ok && e.Err != nil {
		return ErrorMessage(e.Err)
	}
	return "No message provided"
}
