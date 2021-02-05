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
