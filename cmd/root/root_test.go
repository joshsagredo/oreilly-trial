package root

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

//type promptMock struct {
//	// t is not required for this test, but it is would be helpful to assert input parameters if we have it in Run()
//	t *testing.T
//}
//
//func (p promptMock) Run() (string, error) {
//	// return expected result
//	return "", nil
//}
//
//type selectMock struct {
//	t *testing.T
//}
//
//func (p selectMock) Run() (int, string, error) {
//	// return expected result
//	return 1, "No thanks!", nil
//}

//func TestExecuteSelect(t *testing.T) {
//	selectRunner = selectMock{}
//
//	mail.PredefinedValidDomains = []string{"ssss.com"}
//	err := rootCmd.Execute()
//	assert.NotNil(t, err)
//}

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
