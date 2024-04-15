package coderr

import (
	"fmt"
	"os"
)

// Fatal stops program execution
func Fatal(data ...interface{}) {
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

func NotNil(val ...interface{}) {
	for _, v := range val {
		if v == nil {
			fmt.Println(caller(2), "fatal: nil value")
			os.Exit(1)
		}
	}
}
