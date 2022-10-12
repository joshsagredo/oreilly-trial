package cmd

import (
	"errors"
	"fmt"
	"github.com/bilalcaliskan/oreilly-trial/internal/mail"
	"github.com/bilalcaliskan/oreilly-trial/internal/oreilly"
	"github.com/bilalcaliskan/oreilly-trial/internal/random"
	"github.com/manifoldco/promptui"
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
	rootCmd.Flags().IntVarP(&opts.AttemptCount, "attemptCount", "", 1,
		"attempt count of how many times oreilly-trial will try to register again after failed attempts")
	rootCmd.Flags().StringVarP(&opts.LogLevel, "logLevel", "", "info", "log level logging library (debug, info, warn, error)")
	rootCmd.Flags().BoolVarP(&opts.InteractiveMode, "interactiveMode", "", true, "boolean param that "+
		"lets you restart the app after all failed attempts")

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

		var generateFunc = func() error {
			var password string
			var err error
			if password, err = random.GeneratePassword(opts.PasswordRandomLength); err != nil {
				logging.GetLogger().Error("unable to generate password", zap.String("error", err.Error()))
				return err
			}

			tempmails, err := mail.GenerateTempMails(opts.AttemptCount)
			if err != nil {
				logging.GetLogger().Error("an error occurred while generating temp mails", zap.String("error", err.Error()))
				return err
			}

			for i, mail := range tempmails {
				err := oreilly.Generate(opts, mail, password)
				if err == nil {
					logging.GetLogger().Info("trial account successfully created", zap.String("email", mail),
						zap.String("password", password), zap.Int("attempt", i+1))
					return nil
				}

				logging.GetLogger().Error("an error occurred while generating user with tempmail", zap.Int("attempt", i+1),
					zap.String("mail", mail), zap.String("error", err.Error()))
			}

			return errors.New("all attempts are failed, please try to increase attempt count with --attemptCount flag")
		}

		err := generateFunc()
		if err != nil && opts.InteractiveMode {
			fmt.Println()
			prompt := promptui.Prompt{
				Label:     "Would you like to try again?",
				IsConfirm: true,
			}

			for err != nil {
				promptResult, err := prompt.Run()
				if err != nil && strings.ToLower(promptResult) == "n" {
					break
				}

				if strings.ToLower(promptResult) == "y" {
					err = generateFunc()
				}
			}
		}
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
