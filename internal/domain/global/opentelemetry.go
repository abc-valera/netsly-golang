package global

import (
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"

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

func newResource(serviceName string) (*resource.Resource, error) {
	res, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
		),
	)
	if err != nil {
		return nil, coderr.NewInternalErr(err)
	}

	return res, nil
}
