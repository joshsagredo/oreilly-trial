//go:build unit
// +build unit

package mail

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidEmail(t *testing.T) {
	cases := []struct {
		caseName, email string
		shouldPass      bool
	}{
		{"successCase", "nonexistedemail@gmail.com", true},
		{"failCase", "nonexistedmail@gmail", false},
	}

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			t.Logf("running case %s\n", tc.caseName)
			valid := IsValidEmail(tc.email)
			switch tc.shouldPass {
			case true:
				assert.True(t, valid)
			case false:
				assert.False(t, valid)
			}
		})
	}
}
