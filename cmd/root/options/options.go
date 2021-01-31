package options

import "github.com/spf13/cobra"

var rootOptions = &RootOptions{}

// RootOptions contains frequent command line and application options.
type RootOptions struct {
	// BannerFilePath is the relative path to the banner file
	BannerFilePath string
	// VerboseLog is the verbosity of the logging library
	VerboseLog bool
}

func (opts *RootOptions) InitFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&opts.BannerFilePath, "bannerFilePath", "", "banner.txt",
		"relative path of the banner file")
	cmd.Flags().BoolVarP(&opts.VerboseLog, "verbose", "", false,
		"verbose output of the logging library as 'debug' (default false)")
}

// GetRootOptions returns the pointer of RootOptions
func GetRootOptions() *RootOptions {
	return rootOptions
}
