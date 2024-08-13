package infrastructure

import (
	"os"

	"github.com/abc-valera/netsly-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-golang/internal/core/env"
	"github.com/abc-valera/netsly-golang/internal/domain"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/services/emailSender/emailSenderDummy"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/services/logger/loggerNop"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/services/logger/loggerSlog"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/services/opentelemetry/otelHttpExporter"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/services/opentelemetry/otelNopExporter"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/services/opentelemetry/otelWriterExporter"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/services/passworder"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/services/taskQueuer/taskQueuerDummy"
)

func NewServices() domain.Services {
	var services domain.Services

	services.Passworder = passworder.New()

	switch loggerService := env.Load("LOGGER_SERVICE"); loggerService {
	case "nop":
		services.Logger = loggerNop.New()
	case "slogLogger":
		services.Logger = loggerSlog.New(env.Load("SLOG_LOGGER_FOLDER_PATH"))
	default:
		coderr.Fatal("Invalid Logger implementation provided: " + loggerService)
	}

	switch emailSenderService := env.Load("EMAIL_SENDER_SERVICE"); emailSenderService {
	case "dummy":
		services.EmailSender = emailSenderDummy.New()
	default:
		coderr.Fatal("Invalid Email Sender implementation provided: " + emailSenderService)
	}

	switch taskQueuerService := env.Load("TASK_QUEUER_SERVICE"); taskQueuerService {
	case "dummy":
		services.TaskQueuer = taskQueuerDummy.New(services.EmailSender)
	default:
		coderr.Fatal("Invalid Task Queuer implementation provided: " + taskQueuerService)
	}

	// Init OpenTelemetry instrumentation
	switch otelTraceExporterService := env.Load("OTEL_TRACE_EXPORTER"); otelTraceExporterService {
	case "nop":
		services.OtelTraceExporter = coderr.Must(otelNopExporter.New())
	case "console":
		services.OtelTraceExporter = coderr.Must(otelWriterExporter.New(os.Stdout))
	case "tempo":
		services.OtelTraceExporter = coderr.Must(otelHttpExporter.NewTrace(env.Load("TRACE_TEMPO_ENDPOINT")))
	default:
		coderr.Fatal("Invalid Trace Exporter implementation provided: " + otelTraceExporterService)
	}

	// Check if all services are not nil
	coderr.NoErr(coderr.CheckIfStructHasEmptyFields(services))

	return services
}
