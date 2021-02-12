package main

import (
	"os"

	"github.com/ONSDigital/spp-logger/go/spp_logger"
)

func main() {
	context, _ := spp_logger.NewContext("INFO", "uuid.NewString()")
	config := spp_logger.Config{
		Service:     "test_service",
		Component:   "test_component",
		Environment: "test_environment",
		Deployment:  "test_deployment",
		Timezone:    "UTC",
	}
	logger, _ := spp_logger.NewLogger(config, context, "WARNING", os.Stdout)

	logger.Debug("This debug message should not be visible")

	logger.Info("Got to love an info message")

	logger.Warning("But be careful, there be dragons!")

	logger.Error("Error Log")

	logger.Critical("Critical log")

}
