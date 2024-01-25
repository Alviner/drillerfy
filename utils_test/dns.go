package utils_test

import (
	"os"
	"testing"
)

func PostgresDNS(t *testing.T) string {
	t.Helper()

	value, ok := os.LookupEnv("TEST_PG_DNS")
	if !ok {
		return "postgresql://pguser:pgpass@localhost:5432/pgdb"
	}
	return value
}
