package coderr

import (
	"fmt"
	"os"
)

// Must panics if err is not nil and stops program execution
func Must[T any](val T, err error) T {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			os.Exit(1)
		}
	}()
	if err != nil {
		panic(err)
	}
	return val
}

// MustErr panics if err is not nil and stops program execution
func MustErr(err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			os.Exit(1)
		}
	}()
	if err != nil {
		panic(err)
	}
}
