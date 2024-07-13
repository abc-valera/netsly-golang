package middleware

import (
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func NewTracer(tracer trace.Tracer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				instrumentedCtx, span := tracer.Start(
					r.Context(),
					"router",
					trace.WithSpanKind(trace.SpanKindServer),
					trace.WithAttributes(
						attribute.KeyValue{
							Key:   "http.method",
							Value: attribute.StringValue(r.Method),
						},
						attribute.KeyValue{
							Key:   "http.url",
							Value: attribute.StringValue(r.URL.String()),
						},
					),
				)
				defer span.End()

				// Call the next middleware/handler in the chain
				next.ServeHTTP(w, r.WithContext(instrumentedCtx))
			},
		)
	}
}
