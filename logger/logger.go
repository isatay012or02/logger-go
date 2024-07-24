package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Logger struct {
	*logrus.Logger
}

func NewLogger(level logrus.Level, format logrus.Formatter) *Logger {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(level)
	logger.SetFormatter(format)

	return &Logger{logger}
}

func (l *Logger) SetOutputFile(filePath string) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	l.SetOutput(file)

	return nil
}
