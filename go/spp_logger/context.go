package spp_logger

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/sirupsen/logrus"
)

type Context struct {
	logLevel      string
	correlationID string
}

func (context *Context) LogLevel() string {
	return context.logLevel
}

func (context *Context) CorrelationID() string {
	return context.correlationID
}

func (context *Context) IsValid() error {
	if context.logLevel == "" || context.correlationID == "" {
		return fmt.Errorf("Context field missing, must set `logLevel` and `correlationID`")
	}
	if !ValidLevel(context.logLevel) {
		return fmt.Errorf("Log level is not valid, should be one of '%v'", AllLevels)
	}
	return nil
}

func NewContext(logLevel, correlationID string) (*Context, error) {
	var context *Context
	if logLevel == "" && correlationID == "" {
		context = &Context{logLevel: "INFO", correlationID: uuid.NewString()}
	} else if correlationID == "" {
		context = &Context{logLevel: logLevel, correlationID: uuid.NewString()}
	} else {
		context = &Context{logLevel: logLevel, correlationID: correlationID}
	}
	if err := context.IsValid(); err != nil {
		return nil, err
	}
	return context, nil
}

func (context *Context) Fire(entry *logrus.Entry) error {
	fields := logrus.Fields{
		"correlation_id":       context.CorrelationID(),
		"configured_log_level": context.LogLevel(),
	}
	addFieldsToEntry(fields, entry)
	return nil
}

func (context *Context) Levels() []logrus.Level {
	return logrus.AllLevels
}
