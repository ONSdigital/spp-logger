package spp_logger_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"

	"github.com/ONSDigital/spp-logger/go/spp_logger"
)

var _ = Describe("Writing Levels", func() {
	DescribeTable("#WriteLevel",
		func(level logrus.Level, expected string) {
			Expect(spp_logger.WriteLevel(level)).To(Equal(expected))
		},
		Entry("trace should be DEBUG", logrus.TraceLevel, "DEBUG"),
		Entry("debug should be DEBUG", logrus.DebugLevel, "DEBUG"),
		Entry("info should be INFO", logrus.InfoLevel, "INFO"),
		Entry("warning should be WARNING", logrus.WarnLevel, "WARNING"),
		Entry("error should be ERROR", logrus.ErrorLevel, "ERROR"),
		Entry("fatal should be CRITICAL", logrus.FatalLevel, "CRITICAL"),
		Entry("panic should be CRITICAL", logrus.PanicLevel, "CRITICAL"),
		Entry("unknown levels are handled", logrus.Level(2413), "UNKNOWN"),
	)
})

var _ = Describe("Loading Levels", func() {
	DescribeTable("#LoadLevel",
		func(level string, expected logrus.Level) {
			Expect(spp_logger.LoadLevel(level)).To(Equal(expected))
		},
		Entry("DEBUG should be trace", "DEBUG", logrus.TraceLevel),
		Entry("INFO should be info", "INFO", logrus.InfoLevel),
		Entry("WARNING should be warning", "WARNING", logrus.WarnLevel),
		Entry("ERROR should be error", "ERROR", logrus.ErrorLevel),
		Entry("CRITICAL should be fatal", "CRITICAL", logrus.FatalLevel),
		Entry("It should handle any 'case' of level", "cRiticAl", logrus.FatalLevel),
	)
})
