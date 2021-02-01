package spp_logger

import (
	"log"
	// log "github.com/sirupsen/logrus"
)

// type SPPHandler struct {
// 	Config SPPLoggerConfig
// 	Context
// 	LogLevel
// 	Stream
// }

// type SPPLogger struct {
// 	Config SPPLoggerConfig
// 	Name   string
// 	Context
// 	LogLevel
// 	Stream
// }

func Handler() bool {
	return true
}

func OurLog() {
	log.Println("log_text")
}
