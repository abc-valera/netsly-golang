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
	file = strings.Split(file, "netsly-golang/")[1]
	return fmt.Sprintf("[%s:%d]", file, line)
}
