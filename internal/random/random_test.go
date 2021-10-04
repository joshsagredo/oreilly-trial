package random

import (
	"testing"
)

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
			t.Logf("username generated. case=%s, username=%s\n", tc.caseName, username)
		})
	}
}

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
			password := GeneratePassword(tc.randomLength)
			t.Logf("password generated. case=%s, password=%s\n", tc.caseName, password)
		})
	}
}

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
			t.Logf("random email domain selected. case=%s, domain=%s\n", tc.caseName, domain)
		})
	}
}
