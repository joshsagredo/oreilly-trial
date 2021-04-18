package main

import (
	_ "github.com/dimiro1/banner/autoload"
	flag "github.com/spf13/pflag"
	"go.uber.org/zap"
	"oreilly-trial/pkg/logging"
	"oreilly-trial/pkg/oreilly"
)

var (
	logger *zap.Logger
	emailDomain string
	length int
)

func init() {
	logger = logging.GetLogger()

	// for more usable domains, check https://temp-mail.org/
	flag.StringVar(&emailDomain, "emailDomain", "jentrix.com", "usable domain for creating trial " +
		"account, it should be a valid domain")
	flag.IntVar(&length, "length", 12, "length of the random generated username and password")
	flag.Parse()
}

func main() {
	err := oreilly.Generate(emailDomain, length)
	if err != nil {
		logger.Fatal("an error occured while generating user", zap.String("error", err.Error()))
	}
}