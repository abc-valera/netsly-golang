package coderr

import (
	"github.com/gofiber/fiber/v2/log"
)

// Must panics if err is not nil and stops program execution
func Must[T any](val T, err error) T {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal("Fatal error occured: ", r)
		}
	}()
	if err != nil {
		panic(err)
	}
	return val
}
