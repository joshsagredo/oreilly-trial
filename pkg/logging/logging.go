package logging

import "go.uber.org/zap"

var (
	logger *zap.Logger
	err    error
)

func init() {
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

// GetLogger returns the shared *zap.Logger
func GetLogger() *zap.Logger {
	return logger
}
