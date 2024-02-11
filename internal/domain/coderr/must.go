package coderr

import (
	"fmt"
	"os"
)

// Must stops program execution if err is not nil
func Must[T any](val T, err error) T {
	if err != nil {
		fmt.Print("Fatal: ")
		fmt.Println(err)
		os.Exit(1)
	}
	return val
}

// NoErr stops program execution if err is not nil
func NoErr(err error) {
	if err != nil {
		fmt.Print("Fatal: ")
		fmt.Println(err)
		os.Exit(1)
	}
}
