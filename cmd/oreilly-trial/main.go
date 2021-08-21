package main

import (
	"io/ioutil"
	"oreilly-trial/internal/logging"
	"oreilly-trial/internal/options"
	"oreilly-trial/internal/oreilly"
	"os"
	"strings"

	"github.com/dimiro1/banner"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

func init() {
	logger = logging.GetLogger()

	bannerBytes, _ := ioutil.ReadFile("banner.txt")
	banner.Init(os.Stdout, true, false, strings.NewReader(string(bannerBytes)))
}

func main() {
	oto := options.GetOreillyTrialOptions()
	err := oreilly.Generate(oto)
	if err != nil {
		logger.Fatal("an error occurred while generating user", zap.String("error", err.Error()))
	}
}
