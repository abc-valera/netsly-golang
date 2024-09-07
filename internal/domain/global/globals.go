package global

import (
	"context"
	"runtime"
	"strings"
	"sync"

	"go.opentelemetry.io/otel/trace"
)

// global contains global variables (singletons) that are used across the application.

var (
	appModeGlobal   Mode
	appModeInitOnce sync.Once
)

func InitMode(mode Mode) {
	appModeInitOnce.Do(func() {
		appModeGlobal = mode
	})
}

func IsProduction() bool {
	return appModeGlobal == ModeProduction
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
	logGlobal   ILogger
	logInitOnce sync.Once
)

func InitLog(logger ILogger) {
	logInitOnce.Do(func() {
		logGlobal = logger
	})
}

func Log() ILogger {
	return logGlobal
}

var validateGlobal IValidator = newValidator()

func Validate() IValidator {
	return validateGlobal
}
