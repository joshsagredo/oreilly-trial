package main

import (
	"github.com/spf13/pflag"
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
	oto := options.NewOreillyTrialOptions()
	oto.AddFlags(pflag.CommandLine)
	pflag.Parse()
	err := oreilly.Generate(oto)
	if err != nil {
		logger.Fatal("an error occurred while generating user", zap.String("error", err.Error()))
	}
}
