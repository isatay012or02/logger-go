package main

import (
	"github.com/sirupsen/logrus"
	"logger-go/config"
	"logger-go/logger"
)

func main() {

	configuration := config.LoadConfigFromEnv()

	// Создание нового логгера с конфигурацией
	loggergo := logger.NewLogger(configuration.Level, config.GetFormatter(configuration.Format))

	if configuration.OutputFile != "" {
		if err := loggergo.SetOutputFile(configuration.OutputFile); err != nil {
			logrus.Fatalf("Failed to set log output file: %v", err)
		}
	}

	loggergo.Info("This is an info message")
	loggergo.Warn("This is a warning message")
	loggergo.Error("This is an error message")
}
