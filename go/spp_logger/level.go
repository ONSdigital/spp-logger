package spp_logger

import (
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	DebugLevel    = "DEBUG"
	InfoLevel     = "INFO"
	WarningLevel  = "WARNING"
	ErrorLevel    = "ERROR"
	CriticalLevel = "CRITICAL"
)

var AllLevels = []string{
	DebugLevel,
	InfoLevel,
	WarningLevel,
	ErrorLevel,
	CriticalLevel,
}

func WriteLevel(level logrus.Level) string {
	levelMapping := map[logrus.Level]string{
		logrus.TraceLevel: DebugLevel,
		logrus.DebugLevel: DebugLevel,
		logrus.InfoLevel:  InfoLevel,
		logrus.WarnLevel:  WarningLevel,
		logrus.ErrorLevel: ErrorLevel,
		logrus.FatalLevel: CriticalLevel,
		logrus.PanicLevel: CriticalLevel,
	}
	if _, ok := levelMapping[level]; !ok {
		return "UNKNOWN"
	}
	return levelMapping[level]
}

func LoadLevel(level string) logrus.Level {
	levelMapping := map[string]logrus.Level{
		DebugLevel:    logrus.TraceLevel,
		InfoLevel:     logrus.InfoLevel,
		WarningLevel:  logrus.WarnLevel,
		ErrorLevel:    logrus.ErrorLevel,
		CriticalLevel: logrus.FatalLevel,
	}
	return levelMapping[strings.ToUpper(level)]
}

func ValidLevel(level string) bool {
	for _, logLevel := range AllLevels {
		if logLevel == strings.ToUpper(level) {
			return true
		}
	}
	return false
}

type levelHook struct{ CurrentLogLevel string }

func (hook *levelHook) Fire(entry *logrus.Entry) error {
	fields := logrus.Fields{
		"log_level": WriteLevel(entry.Level),
	}

	if hook.CurrentLogLevel == "CRITICAL" {
		fields = logrus.Fields{
			"log_level": hook.CurrentLogLevel,
		}
		hook.CurrentLogLevel = ""
	}

	if _, ok := entry.Data["log_level"]; !ok {
		addFieldsToEntry(fields, entry)
	} else {
		updateEntryFields(fields, entry)
	}
	return nil
}

func (hook *levelHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
