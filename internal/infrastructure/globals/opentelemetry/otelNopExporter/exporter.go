package otelNopExporter

import (
	"io"

	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	traceSDK "go.opentelemetry.io/otel/sdk/trace"
)

func New() (traceSDK.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithWriter(io.Discard),
	)
}
