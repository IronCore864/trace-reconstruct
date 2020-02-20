package config

import (
	"os"
	"strconv"
)

func getEnvIntValue(key string, fallback int) int {
	value, exists := os.LookupEnv(key)
	if exists {
		i, _ := strconv.Atoi(value)
		return i
	}
	return fallback
}

// Config is the struct contains the configuration
type Config struct {
	Interval int
}

func loadConfiguration() Config {
	var conf Config
	conf.Interval = getEnvIntValue("Interval", 20)
	return conf
}

// Conf is the configutation struct object
var Conf = loadConfiguration()
