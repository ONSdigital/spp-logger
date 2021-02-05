package spp_logger_test

import (
	"bytes"

	"time"

	monkey "bou.ke/monkey"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/sirupsen/logrus"

	"github.com/ONSDigital/spp-logger/go/spp_logger"
)

var _ = Describe("the strings package", func() {
	// BeforeEach(func() {
	// 	os.Setenv("SPP_SERVICE", "test_service")
	// 	os.Setenv("SPP_COMPONENT", "test_component")
	// 	os.Setenv("SPP_ENVIRONMENT", "test_env")
	// 	os.Setenv("SPP_DEPLOYMENT", "test_deployment")
	// 	os.Setenv("SPP_USER", "test_user")
	// 	os.Setenv("TIMEZONE", "UTC")
	// })

	// AfterEach(func() {
	// 	os.Clearenv()
	// })

	It("Logs an info message with the correct config and context", func() {
		monkey.Patch(time.Now, func() time.Time {
			return time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
		})

		var buf bytes.Buffer
		context, _ := spp_logger.NewContext("INFO", "test_correlation_id")
		logger, _ := spp_logger.NewLogger(spp_logger.Config{
			Service:     "test_service",
			Component:   "test_component",
			Environment: "test_environment",
			Deployment:  "test_deployment",
			Timezone:    "UTC",
		}, context, logrus.InfoLevel, "INFO", &buf)
		logger.Info("test_message")

		logMessages, err := parseLogLines(buf.String())
		Expect(err).To(BeNil())

		Expect(logMessages[0]["log_level"]).To(Equal("info"))
		Expect(logMessages[0]["timestamp"]).To(Equal("2009-11-17T20:34:58Z"))
		Expect(logMessages[0]["description"]).To(Equal("test_message"))
		Expect(logMessages[0]["correlation_id"]).To(Equal("test_correlation_id"))
		Expect(logMessages[0]["service"]).To(Equal("test_service"))
		Expect(logMessages[0]["component"]).To(Equal("test_component"))
		Expect(logMessages[0]["environment"]).To(Equal("test_environment"))
		Expect(logMessages[0]["deployment"]).To(Equal("test_deployment"))
		Expect(logMessages[0]["timezone"]).To(Equal("UTC"))
	})

	It("Logs a warning message with the correct message", func() {
		var buf bytes.Buffer
		logger, _ := spp_logger.NewLogger(spp_logger.Config{
			Service:     "test_service",
			Component:   "test_component",
			Environment: "test_environment",
			Deployment:  "test_deployment",
			Timezone:    "UTC",
		}, nil, logrus.InfoLevel, "INFO", &buf)
		logger.Warn("test_message")

		logMessages, err := parseLogLines(buf.String())
		Expect(err).To(BeNil())

		Expect(logMessages[0]["log_level"]).To(Equal("warning"))
		Expect(logMessages[0]["description"]).To(Equal("test_message"))

	})

	It("Logs a fatal message with the correct message", func() {
		var buf bytes.Buffer
		logger, _ := spp_logger.NewLogger(spp_logger.Config{
			Service:     "test_service",
			Component:   "test_component",
			Environment: "test_environment",
			Deployment:  "test_deployment",
			Timezone:    "UTC",
		}, nil, logrus.InfoLevel, "INFO", &buf)
		logger.Error("test_message")

		logMessages, err := parseLogLines(buf.String())
		Expect(err).To(BeNil())

		Expect(logMessages[0]["log_level"]).To(Equal("error"))
		Expect(logMessages[0]["description"]).To(Equal("test_message"))

	})

	It("works for multiple logs", func() {
		monkey.Patch(time.Now, func() time.Time {
			return time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
		})

		var buf bytes.Buffer
		logger, _ := spp_logger.NewLogger(spp_logger.Config{
			Service:     "test_service",
			Component:   "test_component",
			Environment: "test_environment",
			Deployment:  "test_deployment",
			Timezone:    "UTC",
		}, nil, logrus.InfoLevel, "INFO", &buf)
		logger.Info("test_message")
		logger.Info("second_test_message")

		logMessagess, err := parseLogLines(buf.String())
		Expect(err).To(BeNil())
		Expect(logMessagess[0]["log_level"]).To(Equal("info"))
	})
})
