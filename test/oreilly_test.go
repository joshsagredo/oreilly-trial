package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"oreilly-trial/internal/options"
	"oreilly-trial/internal/oreilly"
	"testing"
)

var url = "https://learning.oreilly.com/api/v1/registration/individual/"

// this is over real API
// TestGenerateBrokenEmail function tests if Generate function fails with broken arguments and returns desired error
func TestGenerateBrokenEmail(t *testing.T) {
	expectedError := "{\"email\": [\"Enter a valid email address.\"]}"
	oto := options.OreillyTrialOptions{
		CreateUserUrl: url,
		EmailDomains:  []string{"hasan"},
		RandomLength:  12,
	}

	err := oreilly.Generate(&oto)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err.Error())
}

// this is over fake API
// TestGenerateBadRequestResponse function spins up a fake httpserver and simulates a 400 bad request response
func TestGenerateBadRequestResponse(t *testing.T) {
	// Start a local HTTP server
	expectedError := "400 - bad request"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := fmt.Fprint(w, expectedError); err != nil {
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

	err := oreilly.Generate(&oto)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), expectedError)
}

func TestGenerateBrokenAPIUrl(t *testing.T) {
	expectedError := "no such host"
	url := "https://foo.example.com/"
	oto := options.OreillyTrialOptions{
		CreateUserUrl: url,
		EmailDomains:  []string{"jentrix.com"},
		RandomLength:  12,
	}

	err := oreilly.Generate(&oto)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), expectedError)
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
			err := oreilly.Generate(&tc.oto)
			assert.Nil(t, err)
		})
	}
}
