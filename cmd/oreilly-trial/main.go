package main

import (
	"github.com/dimiro1/banner"
	"go.uber.org/zap"
	"io/ioutil"
	"oreilly-trial/internal/logging"
	"oreilly-trial/internal/options"
	"oreilly-trial/internal/oreilly"
	"os"
	"strings"
)

func init() {
	bannerBytes, _ := ioutil.ReadFile("banner.txt")
	banner.Init(os.Stdout, true, false, strings.NewReader(string(bannerBytes)))
}

func main() {
	opts := options.GetOreillyTrialOptions()
	if err := oreilly.Generate(opts); err != nil {
		logging.GetLogger().Fatal("an error occurred while generating user", zap.String("error", err.Error()))
	}
}
