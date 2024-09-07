package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/go-stack/stack"
)

func NewRecoverer() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(rw http.ResponseWriter, r *http.Request) {
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

							if strings.Contains(tFunc, "go.opentelemetry.io") {
								continue
							}

							stackTrace = append(stackTrace, trace)
						}
						panicLocation, _ := strings.CutPrefix(fmt.Sprintf("%+v", stackTrace[1]), "github.com/abc-valera/netsly-golang/")
						global.Log().Error("PANIC_OCCURED",
							"err", err,
							"stack", "["+panicLocation+"]",
						)
					}
				}()
				next.ServeHTTP(rw, r)
			},
		)
	}
}
