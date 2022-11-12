package cmd

import (
	"github.com/bilalcaliskan/oreilly-trial/internal/generator"
	"os"
	"strings"

	"github.com/bilalcaliskan/oreilly-trial/internal/version"

	"github.com/bilalcaliskan/oreilly-trial/internal/logging"
	"github.com/bilalcaliskan/oreilly-trial/internal/options"
	"github.com/dimiro1/banner"
	"github.com/spf13/cobra"
)

var (
	opts *options.OreillyTrialOptions
	ver  = version.Get()
)

func init() {
	opts = options.GetOreillyTrialOptions()
	rootCmd.Flags().StringVarP(&opts.BannerFilePath, "bannerFilePath", "", "build/ci/banner.txt",
		"relative path of the banner file")
	rootCmd.Flags().StringVarP(&opts.LogLevel, "logLevel", "", "info", "log level logging "+
		"library (debug, info, warn, error)")

	if err := rootCmd.Flags().MarkHidden("bannerFilePath"); err != nil {
		logging.GetLogger().Fatalw("fatal error occurred while hiding flag", "error", err.Error())
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
			logging.GetLogger().Errorw("an error occurred while setting log level", "error", err.Error())
			return err
		}

		logging.GetLogger().Infow("oreilly-trial is started", "appVersion", ver.GitVersion,
			"goVersion", ver.GoVersion, "goOS", ver.GoOs, "goArch", ver.GoArch, "gitCommit", ver.GitCommit, "buildDate",
			ver.BuildDate)

		if err := generator.RunGenerator(); err != nil {
			return err
		}

		return nil
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
