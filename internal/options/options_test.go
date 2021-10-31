package options

import "testing"

// TestGetOreillyTrialOptions function tests if GetOreillyTrialOptions function running properly
func TestGetOreillyTrialOptions(t *testing.T) {
	t.Log("fetching default options.OreillyTrialOptions")
	opts := GetOreillyTrialOptions()
	t.Logf("fetched default options.OreillyTrialOptions, %v\n", opts)
}
