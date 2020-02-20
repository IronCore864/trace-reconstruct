package config

import (
	"os"
	"strconv"
	"testing"
)

func setupTestCase(t *testing.T) func(t *testing.T) {
	os.Clearenv()
	return func(t *testing.T) {
		os.Clearenv()
	}
}
func TestGetEnvIntValueDefault(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	env := "Interval"
	defaultInterval := 20
	interval := getEnvIntValue(env, defaultInterval)
	if interval != defaultInterval {
		t.Errorf("Get env var was incorrect, got: %d, want: %d.", interval, defaultInterval)
	}
}

func TestGetEnvIntValueNoneDefault(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	env := "Interval"
	specifiedInterval := 10
	defaultInterval := 20
	os.Setenv(env, strconv.Itoa(specifiedInterval))
	interval := getEnvIntValue(env, defaultInterval)
	if interval != specifiedInterval {
		t.Errorf("Get env var was incorrect, got: %d, want: %d.", interval, specifiedInterval)
	}
}
