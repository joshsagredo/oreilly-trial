package cmd

import (
	"github.com/bilalcaliskan/oreilly-trial/internal/mail"
	"github.com/bilalcaliskan/oreilly-trial/internal/oreilly"
	"github.com/bilalcaliskan/oreilly-trial/internal/random"
	"os"
	"strings"

	"github.com/bilalcaliskan/oreilly-trial/internal/version"

	"github.com/bilalcaliskan/oreilly-trial/internal/logging"
	"github.com/bilalcaliskan/oreilly-trial/internal/options"
	"github.com/dimiro1/banner"
	"github.com/spf13/cobra"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
)

var (
	opts *options.OreillyTrialOptions
	ver  = version.Get()
)

func init() {
	opts = options.GetOreillyTrialOptions()
	rootCmd.Flags().StringVarP(&opts.CreateUserUrl, "createUserUrl", "",
		"https://learning.oreilly.com/api/v1/registration/individual/", "url of the user creation on Oreilly API")
	rootCmd.Flags().IntVarP(&opts.PasswordRandomLength, "passwordRandomLength", "", 16,
		"length of the random generated password between 0 and 32")
	rootCmd.Flags().StringVarP(&opts.BannerFilePath, "bannerFilePath", "", "build/ci/banner.txt",
		"relative path of the banner file")
	rootCmd.Flags().IntVarP(&opts.AttemptCount, "attemptCount", "", 15,
		"attempt count of how many times oreilly-trial will try to register again after failed attempts")
	rootCmd.Flags().StringVarP(&opts.LogLevel, "logLevel", "", "info", "log level logging library (debug, info, warn, error)")

	if err := rootCmd.Flags().MarkHidden("bannerFilePath"); err != nil {
		logging.GetLogger().Fatal("fatal error occured while hiding flag", zap.Error(err))
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "oreilly-trial",
	Short:   "Trial account generator tool for Oreilly",
	Version: ver.GitVersion,
	Long: `As you know, you can create 10 day free trial for https://learning.oreilly.com/ for testing purposes.
This tool does couple of simple steps to provide free trial account for you`,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(opts.BannerFilePath); err == nil {
			bannerBytes, _ := os.ReadFile(opts.BannerFilePath)
			banner.Init(os.Stdout, true, false, strings.NewReader(string(bannerBytes)))
		}

		if err := logging.SetLogLevel(opts.LogLevel); err != nil {
			logging.GetLogger().Error("an error occured while setting log level", zap.Error(err))
			return
		}

		logging.GetLogger().Info("oreilly-trial is started",
			zap.String("appVersion", ver.GitVersion),
			zap.String("goVersion", ver.GoVersion),
			zap.String("goOS", ver.GoOs),
			zap.String("goArch", ver.GoArch),
			zap.String("gitCommit", ver.GitCommit),
			zap.String("buildDate", ver.BuildDate))

		var password string
		var err error
		if password, err = random.GeneratePassword(opts.PasswordRandomLength); err != nil {
			logging.GetLogger().Error("unable to generate password", zap.String("error", err.Error()))
			return
		}

		tempmails, err := mail.GenerateTempMails(opts.AttemptCount)
		if err != nil {
			logging.GetLogger().Error("an error occurred while generating temp mails", zap.String("error", err.Error()))
			return
		}

		for i, mail := range tempmails {
			err := oreilly.Generate(opts, mail, password)
			if err == nil {
				logging.GetLogger().Info("trial account successfully created", zap.String("email", mail),
					zap.String("password", password), zap.Int("attempt", i+1))
				return
			}

			logging.GetLogger().Error("an error occurred while generating user with tempmail", zap.Int("attempt", i+1),
				zap.String("mail", mail), zap.String("error", err.Error()))
		}

		logging.GetLogger().Error("all attempts are failed, please try to increase attempt count with --attemptCount flag")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
