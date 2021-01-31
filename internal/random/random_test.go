//go:build unit
// +build unit

package random

import (
	"fmt"
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
			password, err := GeneratePassword()
			fmt.Println(password)
			assert.Nil(t, err)
			assert.NotEmpty(t, password)
			assert.Len(t, password, randomLength)
		})
	}
}
