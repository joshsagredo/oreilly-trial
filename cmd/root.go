package cmd

import (
	"os"
	"strings"

	"github.com/bilalcaliskan/oreilly-trial/internal/logging"
	"github.com/bilalcaliskan/oreilly-trial/internal/options"
	"github.com/bilalcaliskan/oreilly-trial/internal/oreilly"
	"github.com/dimiro1/banner"
	"github.com/spf13/cobra"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
)

var (
	opts       *options.OreillyTrialOptions
	GitVersion string
)

func init() {
	opts = options.GetOreillyTrialOptions()
	rootCmd.Flags().StringVarP(&opts.CreateUserUrl, "createUserUrl", "",
		"https://learning.oreilly.com/api/v1/registration/individual/", "url of the user creation on Oreilly API")
	rootCmd.Flags().StringSliceVarP(&opts.EmailDomains, "emailDomains", "",
		[]string{"jentrix.com", "geekale.com", "64ge.com", "frnla.com"},
		"comma separated list of usable domain for creating trial account, it should be a valid domain")
	rootCmd.Flags().IntVarP(&opts.UsernameRandomLength, "usernameRandomLength", "", 16,
		"length of the random generated username between 0 and 32")
	rootCmd.Flags().IntVarP(&opts.PasswordRandomLength, "passwordRandomLength", "", 16,
		"length of the random generated password between 0 and 32")
	rootCmd.Flags().StringVarP(&opts.BannerFilePath, "bannerFilePath", "", "build/ci/banner.txt",
		"relative path of the banner file")
	rootCmd.Flags().BoolVarP(&opts.VerboseLog, "verbose", "v", false, "verbose output of the logging library (default false)")

	if err := rootCmd.Flags().MarkHidden("bannerFilePath"); err != nil {
		logging.GetLogger().Fatal("fatal error occured while hiding flag", zap.Error(err))
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "oreilly-trial",
	Short:   "Trial account generator tool for Oreilly",
	Version: GitVersion,
	Long: `As you know, you can create 10 day free trial for https://learning.oreilly.com/ for testing purposes.
This tool does couple of simple steps to provide free trial account for you`,
	Run: func(cmd *cobra.Command, args []string) {
		if opts.VerboseLog {
			logging.Atomic.SetLevel(zap.DebugLevel)
		}

		if _, err := os.Stat(opts.BannerFilePath); err == nil {
			bannerBytes, _ := os.ReadFile(opts.BannerFilePath)
			banner.Init(os.Stdout, true, false, strings.NewReader(string(bannerBytes)))
		}

		if err := oreilly.Generate(opts); err != nil {
			logging.GetLogger().Fatal("an error occurred while generating user", zap.String("error", err.Error()))
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
