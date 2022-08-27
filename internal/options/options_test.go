package options

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetOreillyTrialOptions function tests if GetOreillyTrialOptions function running properly
func TestGetOreillyTrialOptions(t *testing.T) {
	t.Log("fetching default options.OreillyTrialOptions")
	opts := GetOreillyTrialOptions()
	assert.NotNil(t, opts)
	t.Logf("fetched default options.OreillyTrialOptions, %v\n", opts)
}
