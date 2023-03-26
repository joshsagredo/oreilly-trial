package root

import (
	"github.com/bilalcaliskan/oreilly-trial/cmd/root/options"
	"github.com/rs/zerolog"
	"os"
	"strings"

	"github.com/bilalcaliskan/oreilly-trial/internal/generator"
	"github.com/bilalcaliskan/oreilly-trial/internal/mail"
	"github.com/bilalcaliskan/oreilly-trial/internal/oreilly"
	"github.com/bilalcaliskan/oreilly-trial/internal/random"
	"github.com/manifoldco/promptui"

	"github.com/pkg/errors"

	"github.com/bilalcaliskan/oreilly-trial/internal/version"

	"github.com/bilalcaliskan/oreilly-trial/internal/logging"
	"github.com/dimiro1/banner"
	"github.com/spf13/cobra"
)

var (
	opts   *options.RootOptions
	ver    = version.Get()
	logger zerolog.Logger
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:           "oreilly-trial",
		Short:         "Trial account generator tool for Oreilly",
		Version:       ver.GitVersion,
		SilenceErrors: true,
		Long: `As you know, you can create 10 day free trial for https://learning.oreilly.com/ for testing purposes.
This tool does couple of simple steps to provide free trial account for you`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if _, err := os.Stat(opts.BannerFilePath); err == nil {
				bannerBytes, _ := os.ReadFile(opts.BannerFilePath)
				banner.Init(os.Stdout, true, false, strings.NewReader(string(bannerBytes)))
			}

			if opts.VerboseLog {
				logging.EnableDebugLogging()
			}

			logger = logging.GetLogger()
			logger.Info().Str("appVersion", ver.GitVersion).Str("goVersion", ver.GoVersion).Str("goOS", ver.GoOs).
				Str("goArch", ver.GoArch).Str("gitCommit", ver.GitCommit).Str("buildDate", ver.BuildDate).
				Msg("oreilly-trial is started!")

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generator.RunGenerator(); err != nil {
				prompt := promptui.Select{
					Label: "An error occurred while generating Oreilly account with temporary mail, would you like to provide your own valid email address?",
					Items: []string{"Yes please!", "No thanks!"},
				}
				_, result, _ := prompt.Run()
				switch result {
				case "Yes please!":
					prompt := promptui.Prompt{
						Label: "Your valid email address",
						Validate: func(s string) error {
							if !mail.IsValidEmail(s) {
								return errors.Wrap(err, "no valid email provided by user")
							}

							return nil
						},
					}

					mail, _ := prompt.Run()

					password, err := random.GeneratePassword()
					if err != nil {
						logger.Error().
							Str("error", err.Error()).
							Msg("an error occurred while generating password")
						return err
					}

					if err := oreilly.Generate(mail, password, logger); err != nil {
						logger.Error().
							Str("error", err.Error()).
							Msg("an error occurred while generating user with specific email")
						return err
					}

					logger.Info().
						Str("email", mail).
						Str("password", password).
						Msg("trial account successfully created!")

					return nil
				case "No thanks!":
					return err
				}
			}

			return nil
		},
	}
)

func init() {
	opts = options.GetRootOptions()
	opts.InitFlags(rootCmd)

	if err := rootCmd.Flags().MarkHidden("bannerFilePath"); err != nil {
		logger.Warn().Str("error", err.Error()).Msg("an error occurred while hiding flag, ignoring...")
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
