package opentelemetry

import (
	"github.com/abc-valera/netsly-golang/internal/core/coderr"
	traceSDK "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

func NewTracer(exporter traceSDK.SpanExporter, serviceName string) (trace.Tracer, error) {
	res, err := newResource(serviceName)
	if err != nil {
		return nil, coderr.NewInternalErr(err)
	}

	traceProvider := traceSDK.NewTracerProvider(
		traceSDK.WithBatcher(exporter),
		traceSDK.WithResource(res),
	)

	return traceProvider.Tracer("netsly"), nil
}
