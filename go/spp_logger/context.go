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
	return context["correlation_id"]
}

func (context Context) IsValid() error {
	if context["logLevel"] == "" || context["correlation_id"] == "" {
		return fmt.Errorf("Context field missing, must set `logLevel` and `correlation_id`")
	}
	if !ValidLevel(context["logLevel"]) {
		return fmt.Errorf("Log level is not valid, should be one of '%v'", AllLevels)
	}
	return nil
}

func NewContext(logLevel, correlationID string) (Context, error) {
	var context Context
	if logLevel == "" && correlationID == "" {
		context = Context{"logLevel": "INFO", "correlation_id": uuid.NewString()}
	} else if correlationID == "" {
		context = Context{"logLevel": logLevel, "correlation_id": uuid.NewString()}
	} else {
		context = Context{"logLevel": logLevel, "correlation_id": correlationID}
	}

	if err := context.IsValid(); err != nil {
		return nil, err
	}
	return context, nil
}

func (context Context) Fire(entry *logrus.Entry) error {
	contextKeys := context.Keys()
	if len(contextKeys) > 2 {
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
	} else {
		fields := logrus.Fields{
			"correlation_id":       context.CorrelationID(),
			"configured_log_level": context.LogLevel(),
		}
		if _, ok := entry.Data["correlation_id"]; !ok {
			addFieldsToEntry(fields, entry)
		} else {
			updateEntryFields(fields, entry)
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
