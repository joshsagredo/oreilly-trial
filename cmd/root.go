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
	rootCmd.Flags().BoolVarP(&opts.InteractiveMode, "interactiveMode", "", true, "boolean param that "+
		"lets you restart the app after all failed attempts")

	if err := rootCmd.Flags().MarkHidden("bannerFilePath"); err != nil {
		logging.GetLogger().Fatalw("fatal error occured while hiding flag", "error", err.Error())
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:           "oreilly-trial",
	Short:         "Trial account generator tool for Oreilly",
	Version:       ver.GitVersion,
	SilenceErrors: true,
	Long: `As you know, you can create 10 day free trial for https://learning.oreilly.com/ for testing purposes.
This tool does couple of simple steps to provide free trial account for you`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if _, err := os.Stat(opts.BannerFilePath); err == nil {
			bannerBytes, _ := os.ReadFile(opts.BannerFilePath)
			banner.Init(os.Stdout, true, false, strings.NewReader(string(bannerBytes)))
		}

		if err := logging.SetLogLevel(opts.LogLevel); err != nil {
			logging.GetLogger().Errorw("an error occured while setting log level", "error", err.Error())
			return err
		}

		logging.GetLogger().Infow("oreilly-trial is started", "appVersion", ver.GitVersion,
			"goVersion", ver.GoVersion, "goOS", ver.GoOs, "goArch", ver.GoArch, "gitCommit", ver.GitCommit, "buildDate", ver.BuildDate)

		var generateFunc = func() error {
			var password string
			var err error
			if password, err = random.GeneratePassword(opts.PasswordRandomLength); err != nil {
				logging.GetLogger().Errorw("unable to generate password", "error", err.Error())
				return err
			}

			tempmails, err := mail.GenerateTempMails(opts.AttemptCount)
			if err != nil {
				logging.GetLogger().Errorw("an error occurred while generating temp mails", "error", err.Error())
				return err
			}

			for i, mail := range tempmails {
				err := oreilly.Generate(opts, mail, password)
				if err == nil {
					logging.GetLogger().Infow("trial account successfully created", "email", mail, "password", password,
						"attempt", i+1)
					return nil
				}

				logging.GetLogger().Errorw("an error occurred while generating user with tempmail", "attempt", i+1,
					"mail", mail, "error", err.Error())
			}

			err = errors.New("all attempts are failed, please try to increase attempt count with --attemptCount flag")
			logging.GetLogger().Errorw(err.Error(), "attemptCount", opts.AttemptCount)

			return err
		}

		err := generateFunc()
		if err != nil && opts.InteractiveMode {
			fmt.Println()
			for err != nil {
				prompt := promptui.Select{
					Label: "Would you like to try again?",
					Items: []string{"Yes please!", "No thanks!"},
				}

				_, result, _ := prompt.Run()
				if result == "Yes please!" {
					err = generateFunc()
					continue
				}

				break
			}
		}

		return err
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
