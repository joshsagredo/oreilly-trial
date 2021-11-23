package test

import (
	"github.com/stretchr/testify/assert"
	"oreilly-trial/internal/options"
	"testing"
)

// TestGetOreillyTrialOptions function tests if GetOreillyTrialOptions function running properly
func TestGetOreillyTrialOptions(t *testing.T) {
	t.Log("fetching default options.OreillyTrialOptions")
	opts := options.GetOreillyTrialOptions()
	assert.NotNil(t, opts)
	t.Logf("fetched default options.OreillyTrialOptions, %v\n", opts)
}
