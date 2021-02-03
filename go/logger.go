package spp_logger

import (
	// "log"

	"io"
	"time"

	"github.com/sirupsen/logrus"
)

type SPPLogger struct {
	logrus.Logger
	Config SPPLoggerConfig
	Name   string
	// Context
	LogLevel logrus.Level
	// Stream
}

type SPPLoggerEntry struct {
	logrus.Entry
}

type LogMessage struct {
	log_level   string
	timestamp   time.Time
	description string
	service     string
	component   string
	environment string
	deployment  string
}

func NewLogger(config SPPLoggerConfig, logLevel logrus.Level, output io.Writer) *SPPLoggerEntry {
	sppLogger := &SPPLogger{
		Config: config,
	}
	sppLogger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyLevel: "log_level",
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyMsg:   "description",
		},
	})
	sppLogger.SetOutput(output)
	sppLogger.SetLevel(logLevel)
	logrusEntry := sppLogger.WithFields(logrus.Fields{"service": sppLogger.Config.Service, "component": sppLogger.Config.Component, "environment": sppLogger.Config.Environment, "deployment": sppLogger.Config.Deployment, "timezone": sppLogger.Config.Timezone})
	return &SPPLoggerEntry{*logrusEntry}
}

// func OurLog(log_text string, name string, log_level string) {
// 	// log_info := SPPLogger{
// 	// 	Config:   FromEnv(),
// 	// 	Name:     name,
// 	// 	LogLevel: log_level,
// 	// }
// 	// log_message := LogMessage{
// 	// 	log_level:   log_info.LogLevel,
// 	// 	timestamp:   time.Now(),
// 	// 	description: log_text,
// 	// 	service:     log_info.Config.Service,
// 	// 	component:   log_info.Config.Component,
// 	// 	environment: log_info.Config.Environment,
// 	// 	deployment:  log_info.Config.Deployment,
// 	// }

// 	log.SetFormatter(&log.JSONFormatter{})

// 	// log.SetOutput(os.Stdout)

// 	// log.SetLevel(log.WarnLevel)

// 	log.Info(log_text)
// }
