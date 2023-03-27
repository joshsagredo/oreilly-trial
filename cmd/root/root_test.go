package root

import (
	"errors"
	"strconv"
	"testing"

	"github.com/bilalcaliskan/oreilly-trial/internal/mail"
	"github.com/bilalcaliskan/oreilly-trial/internal/prompt"

	"github.com/stretchr/testify/assert"
)

type promptMock struct {
	msg string
	err error
}

func (p promptMock) Run() (string, error) {
	// return expected result
	return p.msg, p.err
}

type selectMock struct {
	msg string
	err error
}

func (p selectMock) Run() (int, string, error) {
	// return expected result
	return 1, p.msg, p.err
}

func TestExecuteWithPromptsSuccessSelectFailPrompt(t *testing.T) {
	// get original value for valid domains
	predefinedValidDomainsOrg := mail.PredefinedValidDomains

	// override valid domains
	mail.PredefinedValidDomains = []string{"ssss.com"}

	selectRunner = selectMock{msg: "Yes please!", err: nil}
	promptRunner = promptMock{msg: "nonexistedemailaddress@example.com", err: errors.New("dummy error")}
	err := rootCmd.Execute()
	assert.NotNil(t, err)

	// revert valid domains
	mail.PredefinedValidDomains = predefinedValidDomainsOrg
	selectRunner = prompt.GetSelectRunner()
	promptRunner = prompt.GetPromptRunner()
}

func TestExecuteWithPromptsFailSelect(t *testing.T) {
	// get original value for valid domains
	predefinedValidDomainsOrg := mail.PredefinedValidDomains

	// override valid domains
	mail.PredefinedValidDomains = []string{"ssss.com"}

	selectRunner = selectMock{msg: "No thanks!", err: nil}
	err := rootCmd.Execute()
	assert.NotNil(t, err)

	// revert valid domains
	mail.PredefinedValidDomains = predefinedValidDomainsOrg
	selectRunner = prompt.GetSelectRunner()
}

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

func TestExecuteVerbose(t *testing.T) {
	verboseOrig, err := rootCmd.Flags().GetBool("verbose")
	assert.Nil(t, err)
	assert.False(t, verboseOrig)

	err = rootCmd.Flags().Set("verbose", "true")
	assert.Nil(t, err)

	err = rootCmd.Execute()
	assert.Nil(t, err)

	_ = rootCmd.Flags().Set("verbose", strconv.FormatBool(verboseOrig))
}
