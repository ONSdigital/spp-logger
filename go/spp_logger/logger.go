package spp_logger

import (
	"io"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	logrus.Logger
	Config   Config
	Name     string
	context  Context
	LogLevel string
}

type ConfigHook struct {
	Config *Config
}

func NewLogger(config Config, context Context, logLevel string, output io.Writer) (*Logger, error) {
	if context["log_level"] != "" {
		context["logLevel"] = context["log_level"]
	}
	if context == nil {
		context, _ = NewContext(logLevel, "")
	}
	if err := context.IsValid(); err != nil {
		return nil, err
	}

	sppLogger := &Logger{
		Config:  config,
		context: context,
	}

	sppLogger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyLevel: "go_log_level",
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyMsg:   "description",
		},
		TimestampFormat: "2006-01-02T15:04:05+00:00",
	})

	sppLogger.SetOutput(output)
	sppLogger.SetLevel(LoadLevel(context.LogLevel()))
	sppLogger.Hooks = make(logrus.LevelHooks)
	sppLogger.AddHook(&ConfigHook{
		Config: &sppLogger.Config,
	})
	sppLogger.AddHook(&levelHook{})
	sppLogger.AddHook(context)
	return sppLogger, nil
}

func (sppLogger *Logger) setContext(context Context) (*Logger, error) {
	sppLogger.SetLevel(LoadLevel(context.LogLevel()))
	sppLogger.AddHook(context)
	return sppLogger, nil
}

func (sppLogger *Logger) OverrideContext(context Context) *Logger {
	if context["log_level"] != "" {
		context["logLevel"] = context["log_level"]
		delete(context, "log_level")
	}
	mainContext := sppLogger.context
	newContext, err := sppLogger.setContext(context)
	if err != nil {
		sppLogger.setContext(mainContext)
	}
	return newContext
}

func (sppLogger *Logger) SetContextAttribute(attribute, value string) *Logger {
	mainContext := sppLogger.context
	context := sppLogger.context
	context[attribute] = value
	newContext, err := sppLogger.setContext(context)
	if err != nil {
		sppLogger.setContext(mainContext)
	}
	return newContext
}

func (sppLogger *Logger) Critical(args ...interface{}) {
	sppLogger.AddHook(&levelHook{CurrentLogLevel: "CRITICAL"})
	sppLogger.Error(args...)

}

func (sppLogger *Logger) Criticalf(format string, args ...interface{}) {
	sppLogger.AddHook(&levelHook{CurrentLogLevel: "CRITICAL"})
	sppLogger.Errorf(format, args...)
}

func (sppLogger *Logger) CriticalFn(fn logrus.LogFunction) {
	sppLogger.AddHook(&levelHook{CurrentLogLevel: "CRITICAL"})
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

func updateEntryFields(fields logrus.Fields, entry *logrus.Entry) {
	for field, value := range fields {
		entry.Data[field] = value
	}
}

func (sppLogger *Logger) Context() Context {
	return sppLogger.context
}
