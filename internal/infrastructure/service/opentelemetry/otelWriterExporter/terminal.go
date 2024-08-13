package otelWriterExporter

import (
	"io"

	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	traceSDK "go.opentelemetry.io/otel/sdk/trace"
)

func New(w io.Writer) (traceSDK.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithPrettyPrint(),
		stdouttrace.WithWriter(w),
	)
}
