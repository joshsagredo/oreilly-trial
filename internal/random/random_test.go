package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGenerate function tests if Generate function running properly
func TestGenerate(t *testing.T) {
	cases := []struct {
		caseName     string
		randomLength int
		outputType   string
	}{
		{"random10username", 10, TypeUsername},
		{"random20username", 20, TypeUsername},
		{"random10password", 10, TypePassword},
		{"random20password", 20, TypePassword},
	}

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			username, err := Generate(tc.randomLength, tc.outputType)
			assert.Nil(t, err)
			assert.NotEmpty(t, username)
			assert.Len(t, username, tc.randomLength)
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

// TestPickEmail tests if PickEmail function is running properly
func TestGenerateInvalidLength(t *testing.T) {
	cases := []struct {
		caseName     string
		randomLength int
		outputType   string
	}{
		{"case1", 660, TypeUsername},
		{"case2", 770, TypeUsername},
		{"case3", 880, TypePassword},
	}

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			res, err := Generate(tc.randomLength, tc.outputType)
			assert.NotNil(t, err)
			assert.Empty(t, res)
		})
	}
}
