package globals

import (
	"os"

	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"github.com/abc-valera/netsly-golang/internal/domain/util/env"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/globals/logger/loggerNop"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/globals/logger/loggerSlog"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/globals/opentelemetry/otelHttpExporter"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/globals/opentelemetry/otelNopExporter"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/globals/opentelemetry/otelWriterExporter"
	traceSDK "go.opentelemetry.io/otel/sdk/trace"
)

func New(entrypoint string) {
	var mode global.Mode
	switch appMode := env.Load("APP_MODE"); appMode {
	case "production", "prod":
		mode = global.ModeProduction
	case "development", "dev":
		mode = global.ModeDevelopment
	default:
		coderr.Fatal("Invalid APP_MODE provided: " + appMode)
	}

	var logger global.ILogger
	switch loggerService := env.Load("LOGGER_SERVICE"); loggerService {
	case "nop":
		logger = loggerNop.New()
	case "slogLogger":
		logger = loggerSlog.New(env.Load("SLOG_LOGGER_FOLDER_PATH"))
	default:
		coderr.Fatal("Invalid Logger implementation provided: " + loggerService)
	}

	var otelTraceExporter traceSDK.SpanExporter
	switch otelTraceExporterService := env.Load("OTEL_TRACE_EXPORTER"); otelTraceExporterService {
	case "nop":
		otelTraceExporter = coderr.Must(otelNopExporter.New())
	case "console":
		otelTraceExporter = coderr.Must(otelWriterExporter.New(os.Stdout))
	case "tempo":
		otelTraceExporter = coderr.Must(otelHttpExporter.NewTrace(env.Load("TRACE_TEMPO_ENDPOINT")))
	default:
		coderr.Fatal("Invalid Trace Exporter implementation provided: " + otelTraceExporterService)
	}

	global.Init(
		mode,
		coderr.Must(global.NewTracer(otelTraceExporter, "netsly."+entrypoint)),
		logger,
	)
}
