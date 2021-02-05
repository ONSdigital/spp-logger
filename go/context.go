package spp_logger

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type SPPLogContext struct {
	LogLevel      string
	CorrelationID string
}

func SetContext(context SPPLogContext) (SPPLogContext, error) {
	if (context.LogLevel == "") || (context.CorrelationID == "") {
		return SPPLogContext{}, fmt.Errorf("Context field missing")
	}
	return context, nil
}

func (context *SPPLogContext) Fire(entry *logrus.Entry) error {
	fields := logrus.Fields{
		"correlation_id": context.CorrelationID,
	}
	addFieldsToEntry(fields, entry)
	return nil
}

func (context *SPPLogContext) Levels() []logrus.Level {
	return logrus.AllLevels
}
