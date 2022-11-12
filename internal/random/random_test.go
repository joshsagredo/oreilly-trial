package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGeneratePassword function tests if GeneratePassword function running properly
func TestGeneratePassword(t *testing.T) {
	cases := []struct {
		caseName string
	}{
		{"randomusername"},
	}

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			username, err := GeneratePassword()
			assert.Nil(t, err)
			assert.NotEmpty(t, username)
			assert.Len(t, username, randomLength)
		})
	}
}
