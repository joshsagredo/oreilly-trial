package oreilly

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"oreilly-trial/internal/options"
	"testing"
)

var (
	url = "https://learning.oreilly.com/api/v1/registration/individual/"
	domains = []string{"jentrix.com"}
)

// TestGenerateBrokenEmail function tests if Generate function fails with broken arguments and returns desired error
func TestGenerate_WhenBrokenEmail_ShouldReturnError(t *testing.T) {
	expectedError := "{\"email\": [\"Enter a valid email address.\"]}"
	mockServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(400)
		if _, err := fmt.Fprint(writer, expectedError); err != nil {
			t.Fatalf("a fatal error occured while writing response body: %s", err.Error())
		}
	}))
	defer mockServer.Close()

	oto := options.OreillyTrialOptions{
		CreateUserUrl: mockServer.URL,
		EmailDomains:  []string{"hasan"},
		RandomLength:  12,
	}

	err := Generate(&oto)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err.Error())
}

// TestGenerate_ShouldReturnError function spins up a fake httpserver and simulates a 400 bad request response
func TestGenerate_ShouldReturnError(t *testing.T) {
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
		EmailDomains:  domains,
		RandomLength:  12,
	}

	err := Generate(&oto)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), expectedError)
}

func TestGenerate_WhenInvalidHost_ShouldReturnError(t *testing.T) {
	expectedError := "no such host"
	url := "https://foo.example.com/"
	oto := options.OreillyTrialOptions{
		CreateUserUrl: url,
		EmailDomains:  domains,
		RandomLength:  12,
	}

	err := Generate(&oto)
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
			EmailDomains:  domains,
			RandomLength:  12,
		}},
		{"case2", options.OreillyTrialOptions{
			CreateUserUrl: url,
			EmailDomains:  domains,
			RandomLength:  16,
		}},
	}

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			err := Generate(&tc.oto)
			assert.Nil(t, err)
		})
	}
}
