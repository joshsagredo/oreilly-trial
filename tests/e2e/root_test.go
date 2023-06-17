//go:build e2e
// +build e2e

package e2e

import (
	"errors"
	"strconv"
	"testing"

	"github.com/bilalcaliskan/oreilly-trial/cmd/root"

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

	root.SelectRunner = selectMock{msg: "Yes please!", err: nil}
	root.PromptRunner = promptMock{msg: "nonexistedemailaddress@example.com", err: errors.New("dummy error")}
	err := root.RootCmd.Execute()
	assert.NotNil(t, err)

	// revert valid domains
	mail.PredefinedValidDomains = predefinedValidDomainsOrg
	root.SelectRunner = prompt.GetSelectRunner()
	root.PromptRunner = prompt.GetPromptRunner()
}

func TestExecuteWithPromptsFailSelect(t *testing.T) {
	// get original value for valid domains
	predefinedValidDomainsOrg := mail.PredefinedValidDomains

	// override valid domains
	mail.PredefinedValidDomains = []string{"ssss.com"}

	root.SelectRunner = selectMock{msg: "No thanks!", err: nil}
	err := root.RootCmd.Execute()
	assert.NotNil(t, err)

	// revert valid domains
	mail.PredefinedValidDomains = predefinedValidDomainsOrg
	root.SelectRunner = prompt.GetSelectRunner()
}

func TestExecute(t *testing.T) {
	bannerFilePathOrig, _ := root.RootCmd.Flags().GetString("bannerFilePath")
	assert.NotNil(t, bannerFilePathOrig)
	assert.NotEmpty(t, bannerFilePathOrig)

	err := root.RootCmd.Flags().Set("bannerFilePath", "./../build/ci/banner.txt")
	assert.Nil(t, err)

	err = root.RootCmd.Execute()
	assert.Nil(t, err)

	_ = root.RootCmd.Flags().Set("bannerFilePath", bannerFilePathOrig)
}

func TestExecute2(t *testing.T) {
	root.Execute()
}

func TestExecuteMissingBannerFile(t *testing.T) {
	bannerFilePathOrig, _ := root.RootCmd.Flags().GetString("bannerFilePath")
	assert.NotNil(t, bannerFilePathOrig)
	assert.NotEmpty(t, bannerFilePathOrig)

	err := root.RootCmd.Flags().Set("bannerFilePath", "asdfasdfasdf")
	assert.Nil(t, err)

	_ = root.RootCmd.Execute()

	_ = root.RootCmd.Flags().Set("bannerFilePath", bannerFilePathOrig)
}

func TestExecuteVerbose(t *testing.T) {
	verboseOrig, err := root.RootCmd.Flags().GetBool("verbose")
	assert.Nil(t, err)
	assert.False(t, verboseOrig)

	err = root.RootCmd.Flags().Set("verbose", "true")
	assert.Nil(t, err)

	err = root.RootCmd.Execute()
	assert.Nil(t, err)

	_ = root.RootCmd.Flags().Set("verbose", strconv.FormatBool(verboseOrig))
}
