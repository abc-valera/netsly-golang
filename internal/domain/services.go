package domain

import (
	"github.com/abc-valera/netsly-golang/internal/domain/service"
	traceSDK "go.opentelemetry.io/otel/sdk/trace"
)

type Services struct {
	Logger            service.ILogger
	EmailSender       service.IEmailSender
	Passworder        service.IPassworder
	TaskQueuer        service.ITaskQueuer
	OtelTraceExporter traceSDK.SpanExporter
}
