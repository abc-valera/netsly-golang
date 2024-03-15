package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/abc-valera/netsly-api-golang/internal/core/global"
	"github.com/go-stack/stack"
)

func Recoverer(next http.Handler) http.Handler {
	fn := func(rw http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				rw.WriteHeader(http.StatusInternalServerError)

				// Check if the error is of type error
				if _, ok := err.(error); !ok {
					err = fmt.Errorf("%v", err)
				}

				var stackTrace stack.CallStack
				// Get the current stacktrace but trim the runtime and
				// then format the stack trace removing the clutter from it
				for _, trace := range stack.Trace().TrimRuntime() {
					tFunc := trace.Frame().Function

					// We don't want this noise to appear in logs
					if tFunc == "runtime.gopanic" {
						continue
					}

					// This call is made before the code reaching our handlers
					if tFunc == "net/http.HandlerFunc.ServeHTTP" {
						break
					}

					stackTrace = append(stackTrace, trace)
				}

				panicLocation, _ := strings.CutPrefix(fmt.Sprintf("%+v", stackTrace[3]), "github.com/abc-valera/netsly-api-golang/")
				global.Log().Error("PANIC_OCCURED",
					"err", err,
					"stack", "["+panicLocation+"]",
				)
			}
		}()
		next.ServeHTTP(rw, r)
	}
	return http.HandlerFunc(fn)
}
