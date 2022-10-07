package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGeneratePassword function tests if GeneratePassword function running properly
func TestGeneratePassword(t *testing.T) {
	cases := []struct {
		caseName     string
		randomLength int
	}{
		{"random10username", 10},
		{"random20username", 20},
		{"random10password", 10},
		{"random20password", 20},
	}

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			username, err := GeneratePassword(tc.randomLength)
			assert.Nil(t, err)
			assert.NotEmpty(t, username)
			assert.Len(t, username, tc.randomLength)
		})
	}
}

// TestGenerateInvalidLength tests if GeneratePassword function fails as expected
func TestGenerateInvalidLength(t *testing.T) {
	cases := []struct {
		caseName     string
		randomLength int
	}{
		{"case1", 660},
		{"case2", 770},
		{"case3", 880},
	}

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			res, err := GeneratePassword(tc.randomLength)
			assert.NotNil(t, err)
			assert.Empty(t, res)
		})
	}
}
