package options

var oreillyTrialOptions = &OreillyTrialOptions{}

// OreillyTrialOptions contains frequent command line and application options.
type OreillyTrialOptions struct {
	// BannerFilePath is the relative path to the banner file
	BannerFilePath string
	// LogLevel is the level of the logging library
	LogLevel string
}

// GetOreillyTrialOptions returns the pointer of OreillyTrialOptions
func GetOreillyTrialOptions() *OreillyTrialOptions {
	return oreillyTrialOptions
}
