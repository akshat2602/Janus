package logger

import (
	"log"

	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitializeLogger() {
	var err error
	Logger, err = zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	Logger.Info("Logger initialized")
}
