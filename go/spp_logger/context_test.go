package spp_logger_test

import (
	"github.com/ONSDigital/spp-logger/go/spp_logger"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/google/uuid"
)

var _ = Describe("#NewContext", func() {
	Context("When neither log_level or correlation_id are set", func() {
		It("Should generate a new context with the default log level and a uuid for correlation_id", func() {
			context, err := spp_logger.NewContext("", "")
			Expect(err).To(BeNil())
			Expect(context.LogLevel()).To(Equal("INFO"))
			_, err = uuid.Parse(context.CorrelationID())
			Expect(err).To(BeNil())
		})
	})

	Context("When just log_level is set", func() {
		It("Should return an error", func() {
			context, err := spp_logger.NewContext("INFO", "")
			Expect(err).To(MatchError("Context field missing, must set `logLevel` and `correlationID`"))
			Expect(context).To(BeNil())
		})
	})

	Context("When just correlation id is set", func() {
		It("Should return an error", func() {
			context, err := spp_logger.NewContext("", "correlation_id")
			Expect(err).To(MatchError("Context field missing, must set `logLevel` and `correlationID`"))
			Expect(context).To(BeNil())
		})
	})

	Context("When both log_level and correlation_id are set", func() {
		It("Should return a configured context", func() {
			context, err := spp_logger.NewContext("INFO", "correlation_id")
			Expect(err).To(BeNil())
			Expect(context.LogLevel()).To(Equal("INFO"))
			Expect(context.CorrelationID()).To(Equal("correlation_id"))
		})
	})
})
