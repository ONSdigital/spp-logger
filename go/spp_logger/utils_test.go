package spp_logger_test

// import (
// 	"errors"

// 	. "github.com/onsi/ginkgo"
// 	. "github.com/onsi/gomega"

// 	"github.com/ONSDigital/spp-logger/go/spp_logger"
// )

// var _ = Describe("the utils", func() {
// 	It("Converts a context struct to a dictionary", func() {
// 		context, err := spp_logger.NewContext("INFO", "test_id")
// 		response_dict, err := spp_logger.ContextToDict(context)

// 		expected_dict := map[string]string{"logLevel": "INFO", "correlationID": "test_id"}

// 		Expect(err).To(BeNil())
// 		Expect(response_dict).To(Equal(expected_dict))
// 	})

// 	It("Converts a dictionary (map from string to string) type to a context", func() {
// 		dict := map[string]string{"logLevel": "WARNING", "correlationID": "test_id"}
// 		expected_context, err := spp_logger.NewContext("WARNING", "test_id")

// 		response_context, err := spp_logger.DictToContext(dict)

// 		Expect(err).To(BeNil())
// 		Expect(response_context).To(Equal(expected_context))
// 	})

// 	It("Converts an empty dictionary (map from string to string) type to a context", func() {
// 		dict := map[string]string{}
// 		responseContext, err := spp_logger.DictToContext(dict)
// 		expectedErr := errors.New("Invalid Dictionary")

// 		Expect(expectedErr).Should(MatchError(err))
// 		Expect(responseContext).To(BeNil())
// 	})
// })
