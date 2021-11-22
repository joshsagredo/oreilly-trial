package oreilly

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"oreilly-trial/internal/options"
	"strings"
	"testing"
)

var url = "https://learning.oreilly.com/api/v1/registration/individual/"

// TestGenerateBrokenEmail function tests if Generate function fails with broken arguments and returns desired error
func TestGenerateBrokenEmail(t *testing.T) {
	expectedError := "{\"email\": [\"Enter a valid email address.\"]}"
	oto := options.OreillyTrialOptions{
		CreateUserUrl: url,
		EmailDomains:  []string{"hasan"},
		RandomLength:  12,
	}

	if err := Generate(&oto); err != nil && err.Error() != expectedError {
		t.Fatalf("expected error should be: %v, got: %v", expectedError, err.Error())
	}
}

func TestGenerateBrokenAPIUrl(t *testing.T) {
	expectedError := "no such host"
	url := "https://foo.example.com/"
	oto := options.OreillyTrialOptions{
		CreateUserUrl: url,
		EmailDomains:  []string{"jentrix.com"},
		RandomLength:  12,
	}

	if err := Generate(&oto); err != nil && !strings.Contains(err.Error(), expectedError) {
		t.Fatalf("expected error should contain: %v, got: %v", expectedError, err.Error())
	}
}

// TestGenerateValidArgs function tests if Generate function running properly with proper values
func TestGenerateValidArgs(t *testing.T) {
	cases := []struct {
		caseName string
		oto      options.OreillyTrialOptions
	}{
		{"case1", options.OreillyTrialOptions{
			CreateUserUrl: url,
			EmailDomains:  []string{"jentrix.com"},
			RandomLength:  12,
		}},
		{"case2", options.OreillyTrialOptions{
			CreateUserUrl: url,
			EmailDomains:  []string{"geekale.com", "64ge.com"},
			RandomLength:  16,
		}},
	}

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			if err := Generate(&tc.oto); err != nil {
				t.Fatalf("expected error should be: nil, got: %v", err.Error())
			}

			t.Logf("trial account successfully created!")
			t.Logf("%v\n", tc.oto)
		})
	}
}

func TestGenerateBrokenJsonResponse(t *testing.T) {
	expectedResponse := "{\"foo\": \"bar\""
	expectedError := "unexpected end of JSON input"
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, expectedResponse)
		if err != nil {
			t.Fatalf("a fatal error occured while writing response body: %s", err.Error())
		}
	}))

	defer func() {
		server.Close()
	}()

	oto := options.OreillyTrialOptions{
		CreateUserUrl: server.URL,
		EmailDomains:  []string{"jentrix.com"},
		RandomLength:  12,
	}

	if err := Generate(&oto); err != nil && err.Error() != expectedError {
		t.Fatalf("expected error should be: %v, got: %v", err.Error(), err.Error())
	}
}
