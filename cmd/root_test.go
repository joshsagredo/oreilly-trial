package cmd

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestExecute(t *testing.T) {
	interactiveModeOrig, _ := rootCmd.Flags().GetBool("interactiveMode")
	assert.NotNil(t, interactiveModeOrig)
	assert.NotEmpty(t, interactiveModeOrig)

	bannerFilePathOrig, _ := rootCmd.Flags().GetString("bannerFilePath")
	assert.NotNil(t, bannerFilePathOrig)
	assert.NotEmpty(t, bannerFilePathOrig)

	attemptCountOrig, _ := rootCmd.Flags().GetInt("attemptCount")
	assert.NotNil(t, attemptCountOrig)
	assert.NotEmpty(t, attemptCountOrig)

	err := rootCmd.Flags().Set("bannerFilePath", "./../build/ci/banner.txt")
	assert.Nil(t, err)
	err = rootCmd.Flags().Set("interactiveMode", strconv.FormatBool(false))
	assert.Nil(t, err)
	err = rootCmd.Flags().Set("attemptCount", strconv.FormatInt(int64(20), 10))
	assert.Nil(t, err)

	err = rootCmd.Execute()
	assert.Nil(t, err)

	_ = rootCmd.Flags().Set("bannerFilePath", bannerFilePathOrig)
	_ = rootCmd.Flags().Set("interactiveMode", strconv.FormatBool(interactiveModeOrig))
	_ = rootCmd.Flags().Set("attemptCount", strconv.FormatInt(int64(attemptCountOrig), 10))
}

func TestExecuteMissingBannerFile(t *testing.T) {
	bannerFilePathOrig, _ := rootCmd.Flags().GetString("bannerFilePath")
	assert.NotNil(t, bannerFilePathOrig)
	assert.NotEmpty(t, bannerFilePathOrig)

	interactiveModeOrig, _ := rootCmd.Flags().GetBool("interactiveMode")
	assert.NotNil(t, interactiveModeOrig)
	assert.NotEmpty(t, interactiveModeOrig)

	err := rootCmd.Flags().Set("bannerFilePath", "asdfasdfasdf")
	assert.Nil(t, err)
	err = rootCmd.Flags().Set("interactiveMode", strconv.FormatBool(false))
	assert.Nil(t, err)

	_ = rootCmd.Execute()

	_ = rootCmd.Flags().Set("bannerFilePath", bannerFilePathOrig)
	_ = rootCmd.Flags().Set("interactiveMode", strconv.FormatBool(interactiveModeOrig))
}

func TestExecuteWrongLogLevel(t *testing.T) {
	logLevelOrig, _ := rootCmd.Flags().GetString("logLevel")
	assert.NotNil(t, logLevelOrig)
	assert.NotEmpty(t, logLevelOrig)

	interactiveModeOrig, _ := rootCmd.Flags().GetBool("interactiveMode")
	assert.NotNil(t, interactiveModeOrig)
	assert.NotEmpty(t, interactiveModeOrig)

	err := rootCmd.Flags().Set("logLevel", "infoooo")
	assert.Nil(t, err)
	err = rootCmd.Flags().Set("interactiveMode", strconv.FormatBool(false))
	assert.Nil(t, err)

	err = rootCmd.Execute()
	assert.NotNil(t, err)

	_ = rootCmd.Flags().Set("logLevel", logLevelOrig)
	_ = rootCmd.Flags().Set("interactiveMode", strconv.FormatBool(interactiveModeOrig))
}

func TestExecuteWrongPasswordLength(t *testing.T) {
	passwordLengthOrig, _ := rootCmd.Flags().GetInt("passwordRandomLength")
	assert.NotNil(t, passwordLengthOrig)
	assert.NotEmpty(t, passwordLengthOrig)

	interactiveModeOrig, _ := rootCmd.Flags().GetBool("interactiveMode")
	assert.NotNil(t, interactiveModeOrig)
	assert.NotEmpty(t, interactiveModeOrig)

	err := rootCmd.Flags().Set("passwordRandomLength", "444")
	assert.Nil(t, err)
	err = rootCmd.Flags().Set("interactiveMode", strconv.FormatBool(false))
	assert.Nil(t, err)

	err = rootCmd.Execute()
	assert.NotNil(t, err)

	_ = rootCmd.Flags().Set("passwordRandomLength", strconv.FormatInt(int64(passwordLengthOrig), 10))
	_ = rootCmd.Flags().Set("interactiveMode", strconv.FormatBool(interactiveModeOrig))
}

func TestExecuteWrongAttemptCount(t *testing.T) {
	attemptCountOrig, _ := rootCmd.Flags().GetInt("attemptCount")
	assert.NotNil(t, attemptCountOrig)
	assert.NotEmpty(t, attemptCountOrig)

	interactiveModeOrig, _ := rootCmd.Flags().GetBool("interactiveMode")
	assert.NotNil(t, interactiveModeOrig)
	assert.NotEmpty(t, interactiveModeOrig)

	err := rootCmd.Flags().Set("attemptCount", strconv.FormatInt(int64(0), 10))
	assert.Nil(t, err)
	err = rootCmd.Flags().Set("interactiveMode", strconv.FormatBool(false))
	assert.Nil(t, err)

	err = rootCmd.Execute()
	assert.NotNil(t, err)

	_ = rootCmd.Flags().Set("attemptCount", strconv.FormatInt(int64(attemptCountOrig), 10))
	_ = rootCmd.Flags().Set("interactiveMode", strconv.FormatBool(interactiveModeOrig))
}
