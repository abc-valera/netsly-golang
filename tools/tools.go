//go:build tools

// Here all the binaries that are used in the project are improted.
// This is a trick to make sure that the binaries are installed with a correct version.
// Thanks to the `//go:build` tag this file is not included in the final binary.

package tools

import (
	// go air is used for hot reloading the application
	_ "github.com/air-verse/air"

	// linters
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/mgechev/revive"

	// mockery is used for generating mocks for interfaces
	_ "github.com/vektra/mockery/v2"
)
