package coderr

import (
	"fmt"
	"os"
)

// Panic stops program execution
func Panic(data interface{}) {
	fmt.Println(caller(2), "panic:", data)
	os.Exit(1)
}

// Must stops program execution if err is not nil
func Must(err error) {
	if err != nil {
		fmt.Println(caller(2), "panic:", err.Error())
		os.Exit(1)
	}
}

// MustWithVal stops program execution if err is not nil
func MustWithVal[T any](val T, err error) T {
	if err != nil {
		fmt.Println(caller(2), "panic:", err.Error())
		os.Exit(1)
	}
	return val
}
