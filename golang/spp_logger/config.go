package spp_logger

import (
	"os"
)

type Config struct {
	Service     string
	Component   string
	Environment string
	Deployment  string
	// User        string //TODO
	Timezone string //TODO
}

func NewConfigFromEnv() *Config {
	LoggerConfig := &Config{
		Service:     getEnv("SPP_SERVICE", ""),
		Component:   getEnv("SPP_COMPONENT", ""),
		Environment: getEnv("SPP_ENVIRONMENT", ""),
		Deployment:  getEnv("SPP_DEPLOYMENT", ""),
		Timezone:    getEnv("TIMEZONE", "UTC"), //TODO
	}

	return LoggerConfig
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
