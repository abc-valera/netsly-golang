package ogen

//go:generate go run github.com/ogen-go/ogen/cmd/ogen@latest --convenient-errors --package ogen --target . --clean ../../internal/port/http/schema/openapi.yml