package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExecute(t *testing.T) {
	bannerFilePathOrig, _ := rootCmd.Flags().GetString("bannerFilePath")
	assert.NotNil(t, bannerFilePathOrig)
	assert.NotEmpty(t, bannerFilePathOrig)

	err := rootCmd.Flags().Set("bannerFilePath", "./../build/ci/banner.txt")
	assert.Nil(t, err)

	err = rootCmd.Execute()
	assert.Nil(t, err)

	_ = rootCmd.Flags().Set("bannerFilePath", bannerFilePathOrig)
}

func TestExecuteMissingBannerFile(t *testing.T) {
	bannerFilePathOrig, _ := rootCmd.Flags().GetString("bannerFilePath")
	assert.NotNil(t, bannerFilePathOrig)
	assert.NotEmpty(t, bannerFilePathOrig)

	err := rootCmd.Flags().Set("bannerFilePath", "asdfasdfasdf")
	assert.Nil(t, err)

	_ = rootCmd.Execute()

	_ = rootCmd.Flags().Set("bannerFilePath", bannerFilePathOrig)
}

func TestExecuteWrongLogLevel(t *testing.T) {
	logLevelOrig, _ := rootCmd.Flags().GetString("logLevel")
	assert.NotNil(t, logLevelOrig)
	assert.NotEmpty(t, logLevelOrig)

	err := rootCmd.Flags().Set("logLevel", "infoooo")
	assert.Nil(t, err)

	err = rootCmd.Execute()
	assert.NotNil(t, err)

	_ = rootCmd.Flags().Set("logLevel", logLevelOrig)
}
