package main

import (
	"os"

	"github.com/ONSDigital/spp-logger/go/spp_logger"
)

func main() {
	// create a context using NewContext method
	context, _ := spp_logger.NewContext("INFO", "uuid.NewString()")

	// create context using a map[string]string
	contextDebug := map[string]string{"log_level": "DEBUG", "correlation_id": "test_id", "survey": "survey", "period": "period"}
	config := spp_logger.Config{
		Service:     "test_service",
		Component:   "test_component",
		Environment: "test_environment",
		Deployment:  "test_deployment",
		Timezone:    "UTC",
	}

	logger, _ := spp_logger.NewLogger(config, context, "DEBUG", os.Stdout)

	logger.Debug("This debug message ========should not======== be visible")

	logger.Info("Got to love an info message")

	logger.Warning("But be careful, there be dragons!")

	logger.Error("Error Log")

	// Adding an attribute to the context
	logger.SetContextAttribute("survey", "survey")

	logger.Critical("Critical log with a survey field")

	logger.OverrideContext(contextDebug)

	logger.Debug("This debug message ========should======== be visible")
}
