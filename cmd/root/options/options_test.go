package options

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetRootOptions function tests if GetRootOptions function running properly
func TestGetRootOptions(t *testing.T) {
	t.Log("fetching default options.RootOptions")
	opts := GetRootOptions()
	assert.NotNil(t, opts)
	t.Logf("fetched default options.RootOptions, %v\n", opts)
}
