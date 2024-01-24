package middleware

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// TODO: this middleware is not working
func Tracer(next http.Handler) http.Handler {
	fn := func(rw http.ResponseWriter, r *http.Request) {
		operation := r.Method + " " + r.URL.Path

		otelhttp.NewHandler(next, operation).ServeHTTP(rw, r)
	}
	return http.HandlerFunc(fn)
}
