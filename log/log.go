package log

import (
	"go.uber.org/zap"
	"os"
	"strings"
)

var (
	_logger   *zap.SugaredLogger
	zapLogger *zap.Logger
)

func init() {
	switch strings.ToLower(os.Getenv("ICITY_SDK_MODE")) {
	case "prod":
		zapLogger, _ = zap.NewProduction()
	default:
		zapLogger, _ = zap.NewDevelopment()
	}
	_logger = zapLogger.WithOptions(zap.AddCallerSkip(1)).Sugar()
}

func Error(err error) {
	_logger.Errorf("error occurs: %v", err)
}

// Debugf logs messages at DEBUG level.
func Debugf(format string, args ...interface{}) {
	_logger.Debugf(format, args...)
}

// Infof logs messages at INFO level.
func Infof(format string, args ...interface{}) {
	_logger.Infof(format, args...)
}

// Warnf logs messages at WARN level.
func Warnf(format string, args ...interface{}) {
	_logger.Warnf(format, args...)
}

// Errorf logs messages at ERROR level.
func Errorf(format string, args ...interface{}) {
	_logger.Errorf(format, args...)
}

// Fatalf logs messages at FATAL level.
func Fatalf(format string, args ...interface{}) {
	_logger.Errorf(format, args...)
}
