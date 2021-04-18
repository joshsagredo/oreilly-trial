package main

import (
	_ "github.com/dimiro1/banner/autoload"
	"go.uber.org/zap"
	"oreilly-trial/pkg/logging"
	"oreilly-trial/pkg/oreilly"
)

var (
	logger *zap.Logger

)

func init() {
	logger = logging.GetLogger()
}

func main() {
	err := oreilly.Generate()
	if err != nil {
		logger.Fatal("an error occured while generating user", zap.String("error", err.Error()))
	}
}