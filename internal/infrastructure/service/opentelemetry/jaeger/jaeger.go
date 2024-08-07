package jaeger

import (
	"context"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"

	traceSDK "go.opentelemetry.io/otel/sdk/trace"
)

func NewTraceExporter() (traceSDK.SpanExporter, error) {
	return otlptrace.New(
		context.TODO(),
		otlptracehttp.NewClient(
			otlptracehttp.WithInsecure(),
			otlptracehttp.WithEndpoint("localhost:4318"),
			otlptracehttp.WithHeaders(map[string]string{
				"content-type": "application/json",
			}),
		),
	)
}
