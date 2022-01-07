package random

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestGenerateUsername function tests if GenerateUsername function running properly
func TestGenerateUsername(t *testing.T) {
	cases := []struct {
		caseName     string
		randomLength int
	}{
		{"random10", 10},
		{"random20", 20},
	}

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			username := GenerateUsername(tc.randomLength)
			assert.NotEmpty(t, username)
			assert.Len(t, username, tc.randomLength)
		})
	}
}

func TestGenerateInvalidLength(t *testing.T) {
	cases := []struct {
		caseName     string
		randomLength int
		outputType   string
	}{
		{"random64username", 64, TypeUsername},
		{"random64password", 64, TypePassword},
	}

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			res, err := Generate(tc.randomLength, tc.outputType)
			assert.NotNil(t, err)
			assert.Empty(t, res)
		})
	}
}

// TestPickEmail tests if PickEmail function is running properly
func TestPickEmail(t *testing.T) {
	cases := []struct {
		caseName     string
		emailDomains []string
	}{
		{"random10", []string{"jentrix.com", "geekale.com", "64ge.com", "frnla.com"}},
		{"random20", []string{"asdfasdfasdf.com", "dsfsdf.com", "64ge.com", "frnla.com"}},
	}

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			t.Logf("picking random email domain for case=%s\n", tc.caseName)
			domain := PickEmail(tc.emailDomains)
			assert.NotNil(t, domain)
			assert.NotEmpty(t, domain)
		})
	}
}
