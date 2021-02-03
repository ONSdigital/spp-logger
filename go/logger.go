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

type ConfigHook struct {
	Config *SPPLoggerConfig
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

func NewLogger(config SPPLoggerConfig, logLevel logrus.Level, output io.Writer) *SPPLogger {
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
	sppLogger.Hooks = make(logrus.LevelHooks)
	sppLogger.AddHook(&ConfigHook{
		Config: &sppLogger.Config,
	})
	return sppLogger
}

func (hook *ConfigHook) Fire(entry *logrus.Entry) error {
	fields := logrus.Fields{"service": hook.Config.Service}
	addFieldsToEntry(fields, entry)
	return nil
}

func (hook *ConfigHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func addFieldsToEntry(fields logrus.Fields, entry *logrus.Entry) {
	for field, value := range fields {
		if _, ok := entry.Data[field]; !ok {
			entry.Data[field] = value
		}
	}
}
