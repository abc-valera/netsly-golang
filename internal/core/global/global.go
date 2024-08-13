package global

import (
	"context"
	"runtime"
	"strings"
	"sync"

	"github.com/abc-valera/netsly-golang/internal/core/app"
	"github.com/abc-valera/netsly-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-golang/internal/core/env"
	"github.com/abc-valera/netsly-golang/internal/core/validator"
	"github.com/abc-valera/netsly-golang/internal/domain/service"
	"go.opentelemetry.io/otel/trace"
)

// global is a package that contains global variables that are used across the application.

var (
	appModeGlobal   app.Mode
	appModeInitOnce sync.Once
)

func InitMode() {
	appModeInitOnce.Do(func() {
		appModeGlobal = coderr.Must(app.NewMode(env.Load("MODE")))
	})
}

func IsProduction() bool {
	return appModeGlobal == app.ModeProduction
}

var (
	tracerGlobal   trace.Tracer
	tracerInitOnce sync.Once
)

func InitTracer(tracer trace.Tracer) {
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
	logInitOnce.Do(func() {
		logGlobal = logger
	})
}

func Log() service.ILogger {
	return logGlobal
}

var validateGlobal validator.IValidator = validator.New()

func Validate() validator.IValidator {
	return validateGlobal
}
