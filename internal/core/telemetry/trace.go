package telemetry

import (
	"context"
	"io"
	"time"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	traceSDK "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func NewTerminalTraceExporter(w io.Writer) (traceSDK.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithPrettyPrint(),
		stdouttrace.WithWriter(w),
	)
}

func NewJaegerTraceExporter() (traceSDK.SpanExporter, error) {
	return otlptrace.New(
		context.TODO(),
		otlptracehttp.NewClient(
			otlptracehttp.WithEndpoint("localhost:4318"),
			otlptracehttp.WithHeaders(map[string]string{
				"content-type": "application/json",
			}),
			otlptracehttp.WithInsecure(),
		),
	)
}

func NewTraceProvider(traceExporter traceSDK.SpanExporter, serviceName string) *traceSDK.TracerProvider {
	traceProvider := traceSDK.NewTracerProvider(
		traceSDK.WithBatcher(traceExporter, traceSDK.WithBatchTimeout(time.Second)),
		traceSDK.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
		)),
	)
	return traceProvider
}
