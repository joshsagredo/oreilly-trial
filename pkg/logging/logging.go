package logging

import "go.uber.org/zap"

var (
	logger *zap.Logger
	err error
)

func init() {
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

func GetLogger() *zap.Logger {
	return logger
}