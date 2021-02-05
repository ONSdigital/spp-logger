package spp_logger

import (
	"fmt"

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

func NewContext(logLevel, correlationID string) *Context {
	return &Context{logLevel: logLevel, correlationID: correlationID}
}

func SetContext(context *Context) (*Context, error) {
	if (context.LogLevel() == "") || (context.CorrelationID() == "") {
		return nil, fmt.Errorf("Context field missing")
	}
	return context, nil
}

func (context *Context) Fire(entry *logrus.Entry) error {
	fields := logrus.Fields{
		"correlation_id": context.CorrelationID(),
	}
	addFieldsToEntry(fields, entry)
	return nil
}

func (context *Context) Levels() []logrus.Level {
	return logrus.AllLevels
}
