package coderr

import (
	"fmt"
	"os"
)

// Fatal stops program execution
func Fatal(data ...any) {
	fmt.Print(caller(2), "fatal: ")
	fmt.Println(data...)
	os.Exit(1)
}

// NoErr stops program execution if err is not nil
func NoErr(err error) {
	if err != nil {
		fmt.Println(caller(2), "fatal:", err.Error())
		os.Exit(1)
	}
}

// Must stops program execution if err is not nil
func Must[T any](val T, err error) T {
	if err != nil {
		fmt.Println(caller(2), "fatal:", err.Error())
		os.Exit(1)
	}
	return val
}

// NoEmpty stops program execution if the provided value is empty/nil
func NoEmpty[T comparable](val T) T {
	var nullValue T
	if val == nullValue {
		fmt.Println(caller(2), "fatal: empty/nil value")
		os.Exit(1)
	}
	return val
}
