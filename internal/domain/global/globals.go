package global

import (
	"context"
	"runtime"
	"strings"
	"sync"
	"time"

	"go.opentelemetry.io/otel/trace"
)

// global contains global variables (singletons) that are used across the application.

var initOnce sync.Once

func Init(
	appMode Mode,
	tracer trace.Tracer,
	log ILogger,

	domainName string,
	subdomainWebApp string,
	subdomainJsonApi string,
) {
	initOnce.Do(func() {
		// Set the timezone to UTC
		time.Local = time.UTC

		// Set the global variables
		appModeGlobal = appMode
		tracerGlobal = tracer
		logGlobal = log
		validateGlobal = newValidator()

		domainNameGlobal = domainName
		subdomainWebAppGlobal = subdomainWebApp
		subdomainJsonApiGlobal = subdomainJsonApi
	})
}

var appModeGlobal Mode

func IsProduction() bool {
	return appModeGlobal == ModeProduction
}

var tracerGlobal trace.Tracer

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

var logGlobal ILogger

func Log() ILogger {
	return logGlobal
}

var validateGlobal IValidator

func Validate() IValidator {
	return validateGlobal
}

var (
	isHttpsGlobal          bool
	domainNameGlobal       string
	subdomainWebAppGlobal  string
	subdomainJsonApiGlobal string
)

func IsHTTPS() bool {
	return isHttpsGlobal
}

func DomainName() string {
	return domainNameGlobal
}

func SubdomainWebApp() string {
	return subdomainWebAppGlobal
}

func SubdomainJsonApi() string {
	return subdomainJsonApiGlobal
}
