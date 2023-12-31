package middleware

import (
	"net/http"
	"time"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/service"
)

// responseWriter is a minimal wrapper for http.ResponseWriter that allows the
// written HTTP status code and body to be captured for logging.
type responseWriter struct {
	http.ResponseWriter

	status int
	body   []byte
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
	}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(data []byte) (int, error) {
	rw.body = data
	return rw.ResponseWriter.Write(data)
}

// NewLoggingMiddleware logs the incoming HTTP request & its duration.
func NewLoggingMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			wrapped := wrapResponseWriter(w)
			next.ServeHTTP(wrapped, r)

			logs := []interface{}{
				"status", wrapped.status,
				"method", r.Method,
				"path", r.URL.EscapedPath(),
				"duration(ms)", time.Since(start).Milliseconds(),
			}
			if wrapped.status != 500 {
				service.Log.Info("REQUEST", logs...)
			} else {
				service.Log.Error("REQUEST", logs...)
			}
		}

		return http.HandlerFunc(fn)
	}
}
