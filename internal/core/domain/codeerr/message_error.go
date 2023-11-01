package codeerr

import (
	"bytes"
	"errors"
	"fmt"
)

type MsgErr struct {
	Code Code   // Code provides general information about the error
	Msg  string // Msg provides additional context in human-readable form
	Err  error  // Err is a nested error
}

func NewMsgErr(code Code, msg string) error {
	return &MsgErr{
		Code: code,
		Msg:  msg,
	}
}

// Error returns the string representation of the Error
func (e *MsgErr) Error() string {
	var buf bytes.Buffer

	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		if e.Code != "" {
			fmt.Fprintf(&buf, "<%s> ", e.Code)
		}
		buf.WriteString(e.Msg)
	}

	return buf.String()
}

// ErrorMessage returns the message of the error, if available. Otherwise returns a error message.
// If error is nil returns an empty string.
func ErrorMessage(err error) string {
	var msgErrTarget *MsgErr
	if err == nil {
		return ""
	} else if e, ok := err.(*MsgErr); ok && e.Msg != "" {
		return e.Msg
	} else if errors.As(err, &msgErrTarget) {
		return msgErrTarget.Msg
	} else if ok && e.Err != nil {
		return ErrorMessage(e.Err)
	}
	return "No message provided"
}
