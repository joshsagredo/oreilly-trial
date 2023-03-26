package root

import (
	"github.com/stretchr/testify/assert"
	"strconv"
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
