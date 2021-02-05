package spp_logger

import (
	// "log"

	"io"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	logrus.Logger
	Config   Config
	Name     string
	Context  *Context
	LogLevel string
}

type ConfigHook struct {
	Config *Config
}

func NewLogger(config Config, context *Context, goLogLevel logrus.Level, logLevel string, output io.Writer) (*Logger, error) {
	if context == nil {
		context, _ = NewContext(logLevel, "correlation_id")
	}
	if err := context.IsValid(); err != nil {
		return nil, err
	}

	sppLogger := &Logger{
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
	sppLogger.AddHook(context)
	return sppLogger, nil
}

func (sppLogger *Logger) Critical(args ...interface{}) {
	sppLogger.Error(args...)
}

func (sppLogger *Logger) Criticalf(format string, args ...interface{}) {
	sppLogger.Errorf(format, args...)
}

func (sppLogger *Logger) CriticalFn(fn logrus.LogFunction) {
	sppLogger.ErrorFn(fn)
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