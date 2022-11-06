package options

var oreillyTrialOptions = &OreillyTrialOptions{}

// OreillyTrialOptions contains frequent command line and application options.
type OreillyTrialOptions struct {
	// CreateUserURL is the url of the user creation on Oreilly API
	CreateUserURL string
	// PasswordRandomLength is the length of the random generated username
	PasswordRandomLength int
	// BannerFilePath is the relative path to the banner file
	BannerFilePath string
	// LogLevel is the level of the logging library
	LogLevel string
}

// GetOreillyTrialOptions returns the pointer of OreillyTrialOptions
func GetOreillyTrialOptions() *OreillyTrialOptions {
	return oreillyTrialOptions
}
