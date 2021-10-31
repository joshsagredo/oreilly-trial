package oreilly

import (
	"oreilly-trial/internal/options"
	"testing"
)

// TestGenerate function tests if Generate function running properly
func TestGenerate(t *testing.T) {
	cases := []struct {
		caseName string
		oto      options.OreillyTrialOptions
	}{
		{"case1", options.OreillyTrialOptions{
			CreateUserUrl: "https://learning.oreilly.com/api/v1/registration/individual/",
			EmailDomains:  []string{"jentrix.com"},
			RandomLength:  12,
		}},
		{"case2", options.OreillyTrialOptions{
			CreateUserUrl: "https://learning.oreilly.com/api/v1/registration/individual/",
			EmailDomains:  []string{"geekale.com", "64ge.com"},
			RandomLength:  16,
		}},
	}

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			if err := Generate(&tc.oto); err != nil {
				t.Fatalf("an error occurred while creating trial user")
			}

			t.Logf("trial account successfully created!")
			t.Logf("%v\n", tc.oto)
		})
	}
}
