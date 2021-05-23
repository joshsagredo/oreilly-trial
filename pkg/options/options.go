package options

import (
	"flag"
	"github.com/spf13/pflag"
)

// NewOreillyTrialOptions returns OreillyTrialOptions pointer with zero values
func NewOreillyTrialOptions() *OreillyTrialOptions {
	return &OreillyTrialOptions{}
}

// OreillyTrialOptions contains frequent command line and application options.
type OreillyTrialOptions struct {
	// CreateUserUrl is the url of the user creation on Oreilly API
	CreateUserUrl string
	// EmailDomains is the comma seperated list of usable domain for creating trial account, it should be a valid domain
	EmailDomains []string
	// RandomLength is the length of the random generated username and password
	RandomLength int
}

func (oto *OreillyTrialOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&oto.CreateUserUrl, "createUserUrl", "https://learning.oreilly.com/api/v1/user/",
		"url of the user creation on Oreilly API")
	// for more usable domains, check https://temp-mail.org/
	fs.StringSliceVar(&oto.EmailDomains, "emailDomains", []string{"jentrix.com", "geekale.com", "64ge.com", "frnla.com"},
		"comma seperated list of usable domain for creating trial account, it should be a valid domain")
	fs.IntVar(&oto.RandomLength, "randomLength", 16, "length of the random generated username and password")
}

func init() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
}
