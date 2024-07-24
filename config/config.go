package config

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	Level      logrus.Level
	Format     string
	OutputFile string
}

func LoadConfigFromEnv() *Config {
	level := logrus.InfoLevel
	if lvl, ok := os.LookupEnv("LOG_LEVEL"); ok {
		if parsedLevel, err := logrus.ParseLevel(lvl); err == nil {
			level = parsedLevel
		}
	}

	format := "json"
	if fmt, ok := os.LookupEnv("LOG_FORMAT"); ok {
		format = fmt
	}

	outputFile := ""
	if file, ok := os.LookupEnv("LOG_OUTPUT_FILE"); ok {
		outputFile = file
	}

	return &Config{
		Level:      level,
		Format:     format,
		OutputFile: outputFile,
	}
}

func GetFormatter(format string) logrus.Formatter {
	switch format {
	case "json":
		return &logrus.JSONFormatter{}
	case "text":
		fallthrough
	default:
		return &logrus.TextFormatter{}
	}
}
