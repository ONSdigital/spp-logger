package main

import (
	"os"

	"github.com/ONSDigital/spp-logger/go/spp_logger"
)

func main() {
	contextInfo, _ := spp_logger.NewContext("INFO", "uuid.NewString()")
	contextDebug, _ := spp_logger.NewContext("DEBUG", "uuid.NewString()")

	config := spp_logger.Config{
		Service:     "test_service",
		Component:   "test_component",
		Environment: "test_environment",
		Deployment:  "test_deployment",
		Timezone:    "UTC",
	}

	logger, _ := spp_logger.NewLogger(config, contextInfo, "WARNING", os.Stdout)

	logger.Debug("This debug message ========should not======== be visible")

	logger.Info("Got to love an info message")

	logger.Warning("But be careful, there be dragons!")

	logger.Error("Error Log")

	logger.Critical("Critical log")

	logger, _ = spp_logger.NewLogger(config, contextDebug, "WARNING", os.Stdout)

	logger.Debug("This debug message ========should======== be visible")

}