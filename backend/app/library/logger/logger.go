package logger

import (
	"log"

	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	var err error

	logger, err = zap.NewDevelopment()
	if err != nil {
		log.Fatal("logger creation error")
	}

	// convert global
	zap.ReplaceGlobals(logger)
}

func Debug(msg string, keysAndValues ...interface{}) {
	defer func() {
		err := logger.Sync()
		if err != nil {
			log.Printf("zap logger sync error -> { level: %s, error: %s}", "debug", err)
		}
	}()
	zap.S().Debugw(msg, keysAndValues...)
}

func Info(msg string, keysAndValues ...interface{}) {
	defer func() {
		err := logger.Sync()
		if err != nil {
			log.Printf("zap logger sync error -> { level: %s, error: %s}", "info", err)
		}
	}()
	zap.S().Infow(msg, keysAndValues...)
}

func Warning(msg string, keysAndValues ...interface{}) {
	defer func() {
		err := logger.Sync()
		if err != nil {
			log.Printf("zap logger sync error -> { level: %s, error: %s}", "warning", err)
		}
	}()
	zap.S().Warnw(msg, keysAndValues...)
}

func Error(msg string, keysAndValues ...interface{}) {
	defer func() {
		err := logger.Sync()
		if err != nil {
			log.Printf("zap logger sync error -> { level: %s, error: %s}", "error", err)
		}
	}()
	zap.S().Errorw(msg, keysAndValues...)
}

func Fatal(msg string, keysAndValues ...interface{}) {
	defer func() {
		err := logger.Sync()
		if err != nil {
			log.Printf("zap logger sync error -> { level: %s, error: %s}", "fatal", err)
		}
	}()
	zap.S().Fatalw(msg, keysAndValues...)
}