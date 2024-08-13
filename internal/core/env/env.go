package env

import (
	"os"
	"strings"
	"time"

	"github.com/abc-valera/netsly-golang/internal/core/coderr"
)

// Load is a shortcut for trimming and empty-cheking environemnt variables.
// If the environment variable is not set, it will exit.
func Load(key string) string {
	env := os.Getenv(key)
	if env == "" {
		coderr.Fatal(key + " environment variable is not set")
	}

	return strings.TrimSpace(env)
}

// LoadDuration is a shortcut for loading and parsing duration from environment variables.
// If the environment variable is not set or parsing error occurs, it will exit.
func LoadDuration(key string) time.Duration {
	env := os.Getenv(key)
	if env == "" {
		coderr.Fatal(key + " environment variable is not set")
	}

	return coderr.Must(time.ParseDuration(env))
}
