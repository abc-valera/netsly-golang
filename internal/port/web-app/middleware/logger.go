package middleware

import (
	"net/http"
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
)

// responseWriter is a minimal wrapper for http.ResponseWriter that allows the
// written HTTP status code and body to be captured for logging
type responseWriter struct {
	http.ResponseWriter

	status int
	body   []byte
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
	}
}

// WriteHeader captures the status code before it is written to the response
func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

// Write captures the response body before it is written to the response
func (rw *responseWriter) Write(data []byte) (int, error) {
	rw.body = data
	return rw.ResponseWriter.Write(data)
}

// Logger logs the incoming HTTP request & its duration
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Wrap the response writer so we can capture the status code and body
			wrapped := newResponseWriter(w)
			// Call the next middleware/handler in the chain
			next.ServeHTTP(wrapped, r)

			// If the status code is not explicitly set, assume 200 OK
			if wrapped.status == 0 {
				wrapped.status = 200
			}

			logMsg := []interface{}{
				"status", wrapped.status,
				"method", r.Method,
				"path", r.URL.EscapedPath(),
				"duration(ms)", time.Since(start).Milliseconds(),
			}
			if wrapped.status < 500 {
				global.Log().Info("REQUEST", logMsg...)
			} else {
				global.Log().Error("REQUEST", logMsg...)
			}
		},
	)
}
