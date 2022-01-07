package cmd

import (
	"github.com/dimiro1/banner"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"io/ioutil"
	"oreilly-trial/internal/logging"
	"oreilly-trial/internal/options"
	"oreilly-trial/internal/oreilly"
	"os"
	"strings"
)

// init initializes the cmd package
func init() {
	opts := options.GetOreillyTrialOptions()
	rootCmd.PersistentFlags().StringVarP(&opts.CreateUserUrl, "createUserUrl", "",
		"https://learning.oreilly.com/api/v1/registration/individual/", "url of the user creation on Oreilly API")
	rootCmd.PersistentFlags().StringSliceVarP(&opts.EmailDomains, "emailDomains", "",
		[]string{"jentrix.com", "geekale.com", "64ge.com", "frnla.com"},
		"comma separated list of usable domain for creating trial account, it should be a valid domain")
	rootCmd.PersistentFlags().IntVarP(&opts.UsernameRandomLength, "usernameRandomLength", "", 16,
		"length of the random generated username between 0 and 32")
	rootCmd.PersistentFlags().IntVarP(&opts.PasswordRandomLength, "passwordRandomLength", "", 16,
		"length of the random generated password between 0 and 32")
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "oreilly-trial",
	Short: "Trial account generator tool for Oreilly",
	Long: `As you know, you can create 10 day free trial for https://learning.oreilly.com/ for testing purposes.
This tool does couple of simple steps to provide free trial account for you`,
	Run: func(cmd *cobra.Command, args []string) {
		opts := options.GetOreillyTrialOptions()

		if err := oreilly.Generate(opts); err != nil {
			logging.GetLogger().Fatal("an error occurred while generating user", zap.String("error", err.Error()))
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	bannerBytes, _ := ioutil.ReadFile("banner.txt")
	banner.Init(os.Stdout, true, false, strings.NewReader(string(bannerBytes)))

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
