package cmd

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestExecute(t *testing.T) {
	attemptCountOrig, _ := rootCmd.Flags().GetInt("attemptCount")
	assert.NotNil(t, attemptCountOrig)
	assert.NotEmpty(t, attemptCountOrig)

	bannerFilePathOrig, _ := rootCmd.Flags().GetString("bannerFilePath")
	assert.NotNil(t, bannerFilePathOrig)
	assert.NotEmpty(t, bannerFilePathOrig)

	err := rootCmd.Flags().Set("bannerFilePath", "./../build/ci/banner.txt")
	assert.Nil(t, err)
	err = rootCmd.Flags().Set("attemptCount", strconv.FormatInt(int64(1), 10))
	assert.Nil(t, err)

	err = rootCmd.Execute()
	assert.Nil(t, err)

	_ = rootCmd.Flags().Set("bannerFilePath", bannerFilePathOrig)
	_ = rootCmd.Flags().Set("attemptCount", strconv.FormatInt(int64(attemptCountOrig), 10))
}

func TestExecuteMissingBannerFile(t *testing.T) {
	bannerFilePathOrig, _ := rootCmd.Flags().GetString("bannerFilePath")
	assert.NotNil(t, bannerFilePathOrig)
	assert.NotEmpty(t, bannerFilePathOrig)
	err := rootCmd.Flags().Set("bannerFilePath", "asdfasdfasdf")
	assert.Nil(t, err)

	err = rootCmd.Execute()
	assert.Nil(t, err)
	_ = rootCmd.Flags().Set("bannerFilePath", bannerFilePathOrig)
}

func TestExecuteWrongLogLevel(t *testing.T) {
	logLevelOrig, _ := rootCmd.Flags().GetString("logLevel")
	assert.NotNil(t, logLevelOrig)
	assert.NotEmpty(t, logLevelOrig)
	err := rootCmd.Flags().Set("logLevel", "infoooo")
	assert.Nil(t, err)

	err = rootCmd.Execute()
	assert.Nil(t, err)

	_ = rootCmd.Flags().Set("logLevel", logLevelOrig)
}

func TestExecuteWrongPasswordLength(t *testing.T) {
	passwordLengthOrig, _ := rootCmd.Flags().GetInt("passwordRandomLength")
	assert.NotNil(t, passwordLengthOrig)
	assert.NotEmpty(t, passwordLengthOrig)

	err := rootCmd.Flags().Set("passwordRandomLength", "444")
	assert.Nil(t, err)

	err = rootCmd.Execute()
	assert.Nil(t, err)

	_ = rootCmd.Flags().Set("passwordRandomLength", strconv.FormatInt(int64(passwordLengthOrig), 10))
}

func TestExecuteWrongAttemptCount(t *testing.T) {
	attemptCountOrig, _ := rootCmd.Flags().GetInt("attemptCount")
	assert.NotNil(t, attemptCountOrig)
	assert.NotEmpty(t, attemptCountOrig)

	err := rootCmd.Flags().Set("attemptCount", strconv.FormatInt(int64(0), 10))
	assert.Nil(t, err)

	err = rootCmd.Execute()
	assert.Nil(t, err)

	_ = rootCmd.Flags().Set("attemptCount", strconv.FormatInt(int64(attemptCountOrig), 10))
}
