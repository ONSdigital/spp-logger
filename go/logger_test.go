package spp_logger_test

import (
	"bytes"
	"log"
	"os"
	"time"

	monkey "bou.ke/monkey"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	// . "github.com/smartystreets/goconvey/convey"
	. "github.com/ONSDigital/spp-logger/go"
)

var _ = Describe("the strings package", func() {
	It("returns `true`", func() {
		Expect(Handler()).To(BeTrue())
	})

	It("logs the string", func() {
		monkey.Patch(time.Now, func() time.Time {
			return time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
		})
		output := captureOutput(func() {
			OurLog()
		})
		Expect(output).To(Equal("2009/11/17 20:34:58 log_text\n"))

	})
})

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}
