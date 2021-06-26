package main

import (
	"go.uber.org/zap"
	"oreilly-trial/pkg/logging"
	"oreilly-trial/pkg/options"
	"oreilly-trial/pkg/oreilly"
)

var (
	logger *zap.Logger
)

func init() {
	logger = logging.GetLogger()
}

func main() {
	oto := options.GetOreillyTrialOptions()
	err := oreilly.Generate(oto)
	if err != nil {
		logger.Fatal("an error occurred while generating user", zap.String("error", err.Error()))
	}
}
