package logwrapper

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger enforces specific log message formats
type Logger struct {
	*logrus.Logger
}

// NewLogger initializes the standard logger
func NewLogger() *Logger {
	var baseLogger = logrus.New()

	var logger = &Logger{baseLogger}

	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetOutput(os.Stdout)
	return logger
}
