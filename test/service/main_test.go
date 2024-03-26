package service_test

import (
	"testing"

	"github.com/abc-valera/netsly-api-golang/test"
	"github.com/stretchr/testify/require"
)

// initTest initializes boilerplate for tests
// and should be called in every test function from this package.
func initTest(t *testing.T) *require.Assertions {
	// Make sure tests are run in parallel
	t.Parallel()

	return require.New(t)
}

func TestMain(m *testing.M) {
	test.InitTestMain()

	m.Run()
}
