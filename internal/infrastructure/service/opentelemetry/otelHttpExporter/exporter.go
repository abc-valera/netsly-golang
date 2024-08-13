package otelHttpExporter

import (
	"context"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	traceSDK "go.opentelemetry.io/otel/sdk/trace"
)

func NewTrace(otlpTraceEndpoint string) (traceSDK.SpanExporter, error) {
	return otlptracehttp.New(context.Background(),
		otlptracehttp.WithInsecure(),
		otlptracehttp.WithEndpoint(otlpTraceEndpoint),
	)
}
