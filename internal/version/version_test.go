package version

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGet(t *testing.T) {
	ver := Get()
	assert.NotNil(t, ver)
	assert.Equal(t, ver.GitVersion, "none")
	assert.Equal(t, ver.GitCommit, "none")
	assert.Equal(t, ver.BuildDate, "none")
}
