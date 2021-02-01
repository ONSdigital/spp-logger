package spp_logger_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	// . "github.com/smartystreets/goconvey/convey"
	. "github.com/ONSDigital/spp-logger/go"
)

// func TestHandler(t *testing.T) {
// 	Convey("Function returns true", t, func() {
// 		response := Handler()
// 		So(response, ShouldBeTrue)
// 	})
// }

var _ = Describe("the strings package", func() {
	It("returns `true`", func() {
		Expect(Handler()).To(BeTrue())
	})
})
