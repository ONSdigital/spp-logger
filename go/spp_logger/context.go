package spp_logger

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/sirupsen/logrus"
)

type Context map[string]string

func (context Context) LogLevel() string {
	return context["logLevel"]
}

func (context Context) CorrelationID() string {
	return context["correlationID"]
}

func (context Context) IsValid() error {
	if context["logLevel"] == "" || context["correlationID"] == "" {
		return fmt.Errorf("Context field missing, must set `logLevel` and `correlationID`")
	}
	if !ValidLevel(context["logLevel"]) {
		return fmt.Errorf("Log level is not valid, should be one of '%v'", AllLevels)
	}
	return nil
}

func NewContext(logLevel, correlationID string) (Context, error) {
	var context Context
	if logLevel == "" && correlationID == "" {
		context = Context{"logLevel": "INFO", "correlationID": uuid.NewString()}
	} else if correlationID == "" {
		context = Context{"logLevel": logLevel, "correlationID": uuid.NewString()}
	} else {
		context = Context{"logLevel": logLevel, "correlationID": correlationID}
	}

	if err := context.IsValid(); err != nil {
		return nil, err
	}
	return context, nil
}

func (context Context) Fire(entry *logrus.Entry) error {
	contextKeys := context.Keys()
	fields := logrus.Fields{
		"correlation_id":       context.CorrelationID(),
		"configured_log_level": context.LogLevel(),
	}
	if _, ok := entry.Data["correlation_id"]; !ok {
		addFieldsToEntry(fields, entry)
	} else {
		updateEntryFields(fields, entry)
	}
	for _, element := range contextKeys {
		field := logrus.Fields{
			element: context[element],
		}
		if _, ok := entry.Data[element]; !ok {
			addFieldsToEntry(field, entry)
		} else {
			updateEntryFields(field, entry)
		}
	}
	return nil
}

func (context Context) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (context Context) Keys() []string {
	keys := make([]string, len(context))
	i := 0
	for k := range context {
		keys[i] = k
		i++
	}

	return keys
}
