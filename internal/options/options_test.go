package options

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestGetOreillyTrialOptions function tests if GetOreillyTrialOptions function running properly
func TestGetOreillyTrialOptions(t *testing.T) {
	t.Log("fetching default options.OreillyTrialOptions")
	opts := GetOreillyTrialOptions()
	assert.NotNil(t, opts)
	t.Logf("fetched default options.OreillyTrialOptions, %v\n", opts)
}
