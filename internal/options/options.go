package options

var oreillyTrialOptions = &OreillyTrialOptions{}

// OreillyTrialOptions contains frequent command line and application options.
type OreillyTrialOptions struct {
	// CreateUserUrl is the url of the user creation on Oreilly API
	CreateUserUrl string
	// PasswordRandomLength is the length of the random generated username
	PasswordRandomLength int
	// BannerFilePath is the relative path to the banner file
	BannerFilePath string
	// AttemptCount is the value of how many times oreilly-trial will try to register again after failed attempts
	AttemptCount int
	// LogLevel is the level of the logging library
	LogLevel string
}

// GetOreillyTrialOptions returns the pointer of OreillyTrialOptions
func GetOreillyTrialOptions() *OreillyTrialOptions {
	return oreillyTrialOptions
}
