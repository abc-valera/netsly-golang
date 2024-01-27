package coderr

import (
	"bytes"
	"errors"
	"fmt"
)

const ErrorMessageNotFound = "No message provided"

type Message struct {
	Code    Code   // Code provides general information about the error
	Message string // Message provides additional context in human-readable form
	Err     error  // Err is a nested error
}

func NewMessage(code Code, message string) error {
	return &Message{
		Code:    code,
		Message: message,
	}
}

func NewMessageWithError(code Code, message string, err error) error {
	return &Message{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// Error returns the string representation of the Error
func (e *Message) Error() string {
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
	var messageErrTarget *Message
	if err == nil {
		return ""
	} else if e, ok := err.(*Message); ok && e.Message != "" {
		return e.Message
	} else if errors.As(err, &messageErrTarget) {
		return messageErrTarget.Message
	} else if ok && e.Err != nil {
		return ErrorMessage(e.Err)
	}
	return ErrorMessageNotFound
}
