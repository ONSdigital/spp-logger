package spp_logger_test

import (
	"bytes"
	"encoding/json"
	"fmt"

	// "os"
	"time"

	monkey "bou.ke/monkey"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	// . "github.com/smartystreets/goconvey/convey"
	. "github.com/ONSDigital/spp-logger/go"
	"github.com/sirupsen/logrus"
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

	It("logs the string input", func() {
		monkey.Patch(time.Now, func() time.Time {
			return time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
		})

		var buf bytes.Buffer
		logger := NewLogger(SPPLoggerConfig{
			Service:     "test_service",
			Component:   "test_component",
			Environment: "test_environment",
			Deployment:  "test_deployment",
			Timezone:    "UTC",
		}, logrus.InfoLevel, &buf)
		logger.Info("test_message")
		// OurLog("log_text", "test", "INFO")

		logMessage := make(map[string]string)
		err := json.Unmarshal(buf.Bytes(), &logMessage)
		Expect(err).To(BeNil())
		fmt.Println(logMessage)
		Expect(logMessage["log_level"]).To(Equal("info"))
		Expect(logMessage["timestamp"]).To(Equal("2009-11-17T20:34:58Z"))
		Expect(logMessage["description"]).To(Equal("test_message"))
		Expect(logMessage["service"]).To(Equal("test_service"))
		Expect(logMessage["component"]).To(Equal("test_component"))
		Expect(logMessage["environment"]).To(Equal("test_environment"))
		Expect(logMessage["deployment"]).To(Equal("test_deployment"))
		Expect(logMessage["timezone"]).To(Equal("UTC"))
	})

})
