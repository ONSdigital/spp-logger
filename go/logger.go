package spp_logger

import (
	// "log"

	"io"
	"time"

	"github.com/sirupsen/logrus"
)

type SPPLogger struct {
	logrus.Logger
	Config   SPPLoggerConfig
	Name     string
	Context  SPPLogContext
	LogLevel string
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

func NewLogger(config SPPLoggerConfig, context SPPLogContext, goLogLevel logrus.Level, logLevel string, output io.Writer) *SPPLogger {
	if context == (SPPLogContext{}) {
		context = SPPLogContext{
			LogLevel:      logLevel,
			CorrelationID: "correlation_id",
		}
	}
	context, err := SetContext(context)
	if err != nil {
	}

	sppLogger := &SPPLogger{
		Config:  config,
		Context: context,
	}

	sppLogger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyLevel: "log_level",
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyMsg:   "description",
		},
	})
	sppLogger.SetOutput(output)
	sppLogger.SetLevel(goLogLevel)
	sppLogger.Hooks = make(logrus.LevelHooks)
	sppLogger.AddHook(&ConfigHook{
		Config: &sppLogger.Config,
	})
	sppLogger.AddHook(&context)
	return sppLogger
}

func (sppLogger *SPPLogger) Critical(args ...interface{}) {
	sppLogger.Error(args...)
}

func (hook *ConfigHook) Fire(entry *logrus.Entry) error {
	fields := logrus.Fields{
		"service":     hook.Config.Service,
		"component":   hook.Config.Component,
		"environment": hook.Config.Environment,
		"deployment":  hook.Config.Deployment,
		"timezone":    hook.Config.Timezone,
	}
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
