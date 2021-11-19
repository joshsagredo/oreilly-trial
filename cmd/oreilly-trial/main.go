package main

import (
	"fmt"
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
	fmt.Println()
}

func main() {
	opts := options.GetOreillyTrialOptions()
	if err := oreilly.Generate(opts); err != nil {
		logger.Fatal("an error occurred while generating user", zap.String("error", err.Error()))
	}
}
