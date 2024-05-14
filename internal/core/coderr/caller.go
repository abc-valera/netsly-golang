package coderr

import (
	"fmt"
	"runtime"
	"strings"
)

// caller returns the string representation of the caller function
func caller(skip int) string {
	// Get the caller
	_, file, line, _ := runtime.Caller(skip)

	var msg string

	// Split the file path
	split := strings.Split(file, "bettermill-accra/")
	if len(split) == 1 {
		msg = fmt.Sprintf("[%s:%d]", file, line)
	} else {
		msg = fmt.Sprintf("[%s:%d]", split[1], line)
	}

	return msg
}
