package options

import (
	"testing"

	"github.com/spf13/cobra"

	"github.com/stretchr/testify/assert"
)

// TestGetRootOptions function tests if GetRootOptions function running properly
func TestGetRootOptions(t *testing.T) {
	t.Log("fetching default options.RootOptions")
	opts := GetRootOptions()
	assert.NotNil(t, opts)
	t.Logf("fetched default options.RootOptions, %v\n", opts)
}

func TestRootOptions_InitFlags(t *testing.T) {
	opts := GetRootOptions()
	assert.NotNil(t, opts)
	cmd := &cobra.Command{}
	opts.InitFlags(cmd)
}
