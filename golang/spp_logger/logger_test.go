package spp_logger_test

import (
	"bytes"
	"time"

	monkey "bou.ke/monkey"

	"github.com/ONSDigital/spp-logger/go/spp_logger"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("the strings package", func() {
	It("has a context and configured log level set to INFO, logger logs INFO, WARNING, ERROR and CRITICAL messages", func() {
		monkey.Patch(time.Now, func() time.Time {
			return time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
		})

		var buf bytes.Buffer
		context, _ := spp_logger.NewContext("INFO", "test_log_correlation_id")
		logger, _ := spp_logger.NewLogger(spp_logger.Config{
			Service:     "test_service",
			Component:   "test_component",
			Environment: "test_environment",
			Deployment:  "test_deployment",
			Timezone:    "UTC",
		}, context, "WARNING", &buf)
		logger.Trace("test_message")
		logger.Info("test_message")
		logger.Warning("warning message")
		logger.Error("error message")
		logger.Critical("critical message")
		logger.Error("error message")

		logMessages, err := parseLogLines(buf.String())
		Expect(err).To(BeNil())

		Expect(logMessages[0]["timestamp"]).To(Equal("2009-11-17T20:34:58+00:00"))
		Expect(logMessages[0]["description"]).To(Equal("test_message"))
		Expect(logMessages[0]["log_correlation_id"]).To(Equal("test_log_correlation_id"))
		Expect(logMessages[0]["service"]).To(Equal("test_service"))
		Expect(logMessages[0]["component"]).To(Equal("test_component"))
		Expect(logMessages[0]["environment"]).To(Equal("test_environment"))
		Expect(logMessages[0]["deployment"]).To(Equal("test_deployment"))
		Expect(logMessages[0]["timezone"]).To(Equal("UTC"))

		Expect(logMessages[0]["log_level"]).To(Equal("INFO"))
		Expect(logMessages[0]["go_log_level"]).To(Equal("info"))
		Expect(logMessages[1]["log_level"]).To(Equal("WARNING"))
		Expect(logMessages[1]["go_log_level"]).To(Equal("warning"))
		Expect(logMessages[2]["log_level"]).To(Equal("ERROR"))
		Expect(logMessages[2]["go_log_level"]).To(Equal("error"))
		Expect(logMessages[3]["log_level"]).To(Equal("CRITICAL"))
		Expect(logMessages[3]["go_log_level"]).To(Equal("error"))
		Expect(logMessages[4]["log_level"]).To(Equal("ERROR"))
		Expect(logMessages[4]["go_log_level"]).To(Equal("error"))
	})

	It("Logs an info message with the correct config and no context", func() {
		monkey.Patch(time.Now, func() time.Time {
			return time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
		})

		context, _ := spp_logger.NewContext("", "")
		var buf bytes.Buffer
		logger, _ := spp_logger.NewLogger(spp_logger.Config{
			Service:     "test_service",
			Component:   "test_component",
			Environment: "test_environment",
			Deployment:  "test_deployment",
			Timezone:    "UTC",
		}, context, "INFO", &buf)
		logger.Info("test_message")

		logMessages, err := parseLogLines(buf.String())
		Expect(err).To(BeNil())

		Expect(logMessages[0]["log_level"]).To(Equal("INFO"))
		Expect(logMessages[0]["timestamp"]).To(Equal("2009-11-17T20:34:58+00:00"))
		Expect(logMessages[0]["description"]).To(Equal("test_message"))
		_, err = uuid.Parse(logMessages[0]["log_correlation_id"])
		Expect(err).To(BeNil())
		Expect(logMessages[0]["service"]).To(Equal("test_service"))
		Expect(logMessages[0]["component"]).To(Equal("test_component"))
		Expect(logMessages[0]["environment"]).To(Equal("test_environment"))
		Expect(logMessages[0]["deployment"]).To(Equal("test_deployment"))
		Expect(logMessages[0]["timezone"]).To(Equal("UTC"))
	})

	It("has nil context and configured log level set to DEBUG, logs a debug message with the correct message", func() {
		var buf bytes.Buffer
		logger, _ := spp_logger.NewLogger(spp_logger.Config{
			Service:     "test_service",
			Component:   "test_component",
			Environment: "test_environment",
			Deployment:  "test_deployment",
			Timezone:    "UTC",
		}, nil, "DEBUG", &buf)
		logger.Debug("test_message")

		logMessages, err := parseLogLines(buf.String())
		Expect(err).To(BeNil())

		Expect(logMessages[0]["log_level"]).To(Equal("DEBUG"))
		Expect(logMessages[0]["go_log_level"]).To(Equal("debug"))
		Expect(logMessages[0]["description"]).To(Equal("test_message"))
	})

	It("has an extended context and configured log level set to INFO, logs an INFO message with the correct message", func() {
		var buf bytes.Buffer
		context := map[string]string{"logLevel": "INFO", "log_correlation_id": "test_id", "survey": "survey", "period": "period"}
		logger, _ := spp_logger.NewLogger(spp_logger.Config{
			Service:     "test_service",
			Component:   "test_component",
			Environment: "test_environment",
			Deployment:  "test_deployment",
			Timezone:    "UTC",
		}, context, "DEBUG", &buf)
		logger.Info("test_message")

		logMessages, err := parseLogLines(buf.String())
		Expect(err).To(BeNil())

		Expect(logMessages[0]["log_level"]).To(Equal("INFO"))
		Expect(logMessages[0]["go_log_level"]).To(Equal("info"))
		Expect(logMessages[0]["description"]).To(Equal("test_message"))
		Expect(logMessages[0]["log_correlation_id"]).To(Equal("test_id"))
		Expect(logMessages[0]["survey"]).To(Equal("survey"))
		Expect(logMessages[0]["period"]).To(Equal("period"))
	})

	It("Override method works successfully", func() {
		var buf bytes.Buffer
		context := map[string]string{"logLevel": "INFO", "log_correlation_id": "test_id", "survey": "survey", "period": "period"}
		logger, _ := spp_logger.NewLogger(spp_logger.Config{
			Service:     "test_service",
			Component:   "test_component",
			Environment: "test_environment",
			Deployment:  "test_deployment",
			Timezone:    "UTC",
		}, context, "DEBUG", &buf)
		context, _ = spp_logger.NewContext("DEBUG", "ID")

		logger.Debug("test_message_fail")

		logger.OverrideContext(context)

		logger.Debug("test_message_success")

		logMessages, err := parseLogLines(buf.String())
		Expect(err).To(BeNil())

		Expect(logMessages[0]["log_level"]).To(Equal("DEBUG"))
		Expect(logMessages[0]["go_log_level"]).To(Equal("debug"))
		Expect(logMessages[0]["configured_log_level"]).To(Equal("DEBUG"))

		Expect(logMessages[0]["description"]).To(Equal("test_message_success"))

	})

	It("SetContextAttribute method works successfully", func() {
		var buf bytes.Buffer
		context, _ := spp_logger.NewContext("DEBUG", "ID")
		logger, _ := spp_logger.NewLogger(spp_logger.Config{
			Service:     "test_service",
			Component:   "test_component",
			Environment: "test_environment",
			Deployment:  "test_deployment",
			Timezone:    "UTC",
		}, context, "DEBUG", &buf)

		logger.Info("test_info_message")

		logMessages, err := parseLogLines(buf.String())
		Expect(err).To(BeNil())

		Expect(logMessages[0]["survey"]).To(Equal(""))

		logger.SetContextAttribute("survey", "survey")

		logger.Debug("test_message_success")

		logMessages, err = parseLogLines(buf.String())
		Expect(err).To(BeNil())
		Expect(logMessages[1]["survey"]).To(Equal("survey"))
	})

	It("Takes in a context with `log_level` instead of `logLevel` and still works", func() {
		var buf bytes.Buffer
		context := map[string]string{"logLevel": "WARNING", "log_correlation_id": "test_id"}
		contextInfo := map[string]string{"logLevel": "INFO", "correlation_id": "test_id"}
		contextDebug := map[string]string{"log_level": "DEBUG", "correlation_id": "test_id"}

		logger, _ := spp_logger.NewLogger(spp_logger.Config{
			Service:     "test_service",
			Component:   "test_component",
			Environment: "test_environment",
			Deployment:  "test_deployment",
			Timezone:    "UTC",
		}, context, "CRITICAL", &buf)

		logger.Critical("test_info_message")

		logger.OverrideContext(contextInfo)
		logger.Info("test_info_message")
		logger.OverrideContext(contextDebug)
		logger.Debug("test_debug_message")

		logMessages, err := parseLogLines(buf.String())
		Expect(err).To(BeNil())

		Expect(logMessages[0]["configured_log_level"]).To(Equal("WARNING"))
		Expect(logMessages[1]["configured_log_level"]).To(Equal("INFO"))
		Expect(logMessages[2]["configured_log_level"]).To(Equal("DEBUG"))

	})

	It("Context method returns context", func() {
		var buf bytes.Buffer
		context := map[string]string{"logLevel": "INFO", "log_correlation_id": "test_id"}
		logger, _ := spp_logger.NewLogger(spp_logger.Config{
			Service:     "test_service",
			Component:   "test_component",
			Environment: "test_environment",
			Deployment:  "test_deployment",
			Timezone:    "UTC",
		}, context, "DEBUG", &buf)

		response := logger.Context()
		Expect(response["logLevel"]).To(Equal(context["logLevel"]))
		Expect(response["log_correlation_id"]).To(Equal(context["log_correlation_id"]))
	})

})
