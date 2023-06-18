//go:build e2e
// +build e2e

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

	SelectRunner = selectMock{msg: "Yes please!", err: nil}
	PromptRunner = promptMock{msg: "nonexistedemailaddress@example.com", err: errors.New("dummy error")}
	err := RootCmd.Execute()
	assert.NotNil(t, err)

	// revert valid domains
	mail.PredefinedValidDomains = predefinedValidDomainsOrg
	SelectRunner = prompt.GetSelectRunner()
	PromptRunner = prompt.GetPromptRunner()
}

func TestExecuteWithPromptsFailSelect(t *testing.T) {
	// get original value for valid domains
	predefinedValidDomainsOrg := mail.PredefinedValidDomains

	// override valid domains
	mail.PredefinedValidDomains = []string{"ssss.com"}

	SelectRunner = selectMock{msg: "No thanks!", err: nil}
	err := RootCmd.Execute()
	assert.NotNil(t, err)

	// revert valid domains
	mail.PredefinedValidDomains = predefinedValidDomainsOrg
	SelectRunner = prompt.GetSelectRunner()
}

func TestExecute(t *testing.T) {
	bannerFilePathOrig, _ := RootCmd.Flags().GetString("bannerFilePath")
	assert.NotNil(t, bannerFilePathOrig)
	assert.NotEmpty(t, bannerFilePathOrig)

	err := RootCmd.Flags().Set("bannerFilePath", "./../build/ci/banner.txt")
	assert.Nil(t, err)

	err = RootCmd.Execute()
	assert.Nil(t, err)

	_ = RootCmd.Flags().Set("bannerFilePath", bannerFilePathOrig)
}

func TestExecute2(t *testing.T) {
	Execute()
}

func TestExecuteMissingBannerFile(t *testing.T) {
	bannerFilePathOrig, _ := RootCmd.Flags().GetString("bannerFilePath")
	assert.NotNil(t, bannerFilePathOrig)
	assert.NotEmpty(t, bannerFilePathOrig)

	err := RootCmd.Flags().Set("bannerFilePath", "asdfasdfasdf")
	assert.Nil(t, err)

	_ = RootCmd.Execute()

	_ = RootCmd.Flags().Set("bannerFilePath", bannerFilePathOrig)
}

func TestExecuteVerbose(t *testing.T) {
	verboseOrig, err := RootCmd.Flags().GetBool("verbose")
	assert.Nil(t, err)
	assert.False(t, verboseOrig)

	err = RootCmd.Flags().Set("verbose", "true")
	assert.Nil(t, err)

	err = RootCmd.Execute()
	assert.Nil(t, err)

	_ = RootCmd.Flags().Set("verbose", strconv.FormatBool(verboseOrig))
}
