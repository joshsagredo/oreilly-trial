package options

var oreillyTrialOptions = &OreillyTrialOptions{}

// OreillyTrialOptions contains frequent command line and application options.
type OreillyTrialOptions struct {
	// CreateUserUrl is the url of the user creation on Oreilly API
	CreateUserUrl string
	// EmailDomains is the comma separated list of usable domain for creating trial account, it should be a valid domain
	EmailDomains []string
	// UsernameRandomLength is the length of the random generated username
	UsernameRandomLength int
	// PasswordRandomLength is the length of the random generated username
	PasswordRandomLength int
	// BannerFilePath is the relative path to the banner file
	BannerFilePath string
	// VerboseLog is the verbosity of the logging library
	VerboseLog bool
}

// GetOreillyTrialOptions returns the pointer of OreillyTrialOptions
func GetOreillyTrialOptions() *OreillyTrialOptions {
	return oreillyTrialOptions
}
