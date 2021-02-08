package spp_logger_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ONSDigital/spp-logger/go/spp_logger"
)

var _ = Describe("the utils", func() {
	It("Changes a context type to a dictionary", func() {
		context, err := spp_logger.NewContext("INFO", "test_id")
		response_dict := spp_logger.ContextToDict(context)

		expected_dict := map[string]string{"logLevel": "INFO", "correlationID": "test_id"}

		Expect(err).To(BeNil())
		Expect(response_dict).To(Equal(expected_dict))

	})

	It("Changes a context type to a dictionary", func() {
		dict := map[string]string{"logLevel": "WARNING", "correlationID": "test_id"}
		response_context, err := spp_logger.DictToContext(dict)

		expected_context, err := spp_logger.NewContext("WARNING", "test_id")

		Expect(err).To(BeNil())
		Expect(response_context).To(Equal(expected_context))
	})
})
