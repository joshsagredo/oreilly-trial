package test

import (
	"github.com/stretchr/testify/assert"
	"oreilly-trial/internal/random"
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
			username := random.GenerateUsername(tc.randomLength)
			assert.NotEmpty(t, username)
			assert.Len(t, username, tc.randomLength)
		})
	}
}

// TestGeneratePassword function tests if GeneratePassword function running properly
func TestGeneratePassword(t *testing.T) {
	cases := []struct {
		caseName     string
		randomLength int
	}{
		{"random10", 10},
		{"random20", 20},
	}

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			password := random.GeneratePassword(tc.randomLength)
			assert.NotEmpty(t, password)
			assert.Len(t, password, tc.randomLength)
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
			domain := random.PickEmail(tc.emailDomains)
			assert.NotNil(t, domain)
			assert.NotEmpty(t, domain)
		})
	}
}
