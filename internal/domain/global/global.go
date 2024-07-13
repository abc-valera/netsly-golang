package global

import (
	"context"
	"runtime"
	"strings"
	"sync"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/core/mode"
	"github.com/abc-valera/netsly-api-golang/internal/core/validator"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
	"go.opentelemetry.io/otel/trace"
)

// global is a package that contains global variables that are used across the application.

var (
	appModeGlobal   mode.Mode
	appModeInitOnce sync.Once
)

func InitMode(appMode mode.Mode) {
	if appMode != mode.Development && appMode != mode.Production {
		coderr.Fatal("Provided invalid application mode. Should be either 'development' or 'production'")
	}

	appModeInitOnce.Do(func() {
		appModeGlobal = appMode
	})
}

func Mode() mode.Mode {
	return appModeGlobal
}

var (
	tracerGlobal   trace.Tracer
	tracerInitOnce sync.Once
)

func InitTracer(tracer trace.Tracer) {
	coderr.NotNil(tracer)

	tracerInitOnce.Do(func() {
		tracerGlobal = tracer
	})
}

func Tracer() trace.Tracer {
	return tracerGlobal
}

// NewSpan is a wrapper for tracer.Start that sets the span name to the calling function's name
func NewSpan(ctx context.Context) (context.Context, trace.Span) {
	// Get the name of the calling function
	pc, _, _, _ := runtime.Caller(1)
	split := strings.Split(runtime.FuncForPC(pc).Name(), "/")
	funcName := split[len(split)-1]

	return tracerGlobal.Start(ctx, funcName)
}

var (
	logGlobal   service.ILogger
	logInitOnce sync.Once
)

func InitLog(logger service.ILogger) {
	coderr.NotNil(logger)

	logInitOnce.Do(func() {
		logGlobal = logger
	})
}

func Log() service.ILogger {
	return logGlobal
}

var validate validator.IValidator = validator.New()

func Validate() validator.IValidator {
	return validate
}
