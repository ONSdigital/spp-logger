package spp_logger

import (
	"os"
)

type SPPLoggerConfig struct {
	Service     string
	Component   string
	Environment string
	Deployment  string
	User        string //TODO
	Timezone    string //TODO
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func FromEnv() SPPLoggerConfig {

	LoggerConfig := SPPLoggerConfig{
		Service:     GetEnv("SPP_SERVICE", ""),
		Component:   GetEnv("SPP_COMPONENT", ""),
		Environment: GetEnv("SPP_ENVIRONMENT", ""),
		Deployment:  GetEnv("SPP_DEPLOYMENT", ""),
		User:        GetEnv("SPP_USER", "nil"), //TODO
		Timezone:    GetEnv("TIMEZONE", "UTC"), //TODO
	}

	return LoggerConfig
}
