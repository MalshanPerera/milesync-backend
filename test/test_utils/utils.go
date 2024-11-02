package test_utils

import (
	"os"
	"testing"
)

func SkipCI(t *testing.T) {
	if os.Getenv("ENV") != "" {
		t.Skip("Skipping testing in CI environment")
	}
}
