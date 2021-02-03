package spp_logger_test

import (
	"os"

	. "github.com/ONSDigital/spp-logger/go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("The config", func() {
	BeforeEach(func() {
		os.Setenv("SPP_SERVICE", "test")
		os.Setenv("SPP_COMPONENT", "test")
		os.Setenv("SPP_ENVIRONMENT", "test")
		os.Setenv("SPP_DEPLOYMENT", "test")
		// os.Setenv("SPP_USER", "test_user")
		os.Setenv("TIMEZONE", "UTC")
	})

	AfterEach(func() {
		os.Clearenv()
	})

	It("returns the correct environment variables set in os", func() {
		expected := SPPLoggerConfig{
			Service:     "test",
			Component:   "test",
			Environment: "test",
			Deployment:  "test",
			// User:        "test_user",
			Timezone: "UTC",
		}
		Expect(FromEnv()).Should(Equal(expected))
	})
})
