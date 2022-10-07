package oreilly

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bilalcaliskan/oreilly-trial/internal/mailslurp"

	"github.com/bilalcaliskan/oreilly-trial/internal/options"
	"github.com/stretchr/testify/assert"
)

var url = "https://learning.oreilly.com/api/v1/registration/individual/"

// TestGenerateError function spins up a fake httpserver and simulates a 400 bad request response
func TestGenerateError(t *testing.T) {
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
		CreateUserUrl:        server.URL,
		PasswordRandomLength: 12,
	}

	err := Generate(&oto)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), expectedError)
}

// TestGenerateInvalidHost function tests if Generate function fails on broken Host argument
func TestGenerateInvalidHost(t *testing.T) {
	expectedError := "no such host"
	url := "https://foo.example.com/"
	oto := options.OreillyTrialOptions{
		CreateUserUrl:        url,
		PasswordRandomLength: 12,
	}

	err := Generate(&oto)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), expectedError)
}

func TestGenerateInvalidRandom(t *testing.T) {
	cases := []struct {
		caseName string
		oto      options.OreillyTrialOptions
	}{
		{"case1", options.OreillyTrialOptions{
			CreateUserUrl:        url,
			PasswordRandomLength: 666,
		}},
		{"case2", options.OreillyTrialOptions{
			CreateUserUrl:        url,
			PasswordRandomLength: 665,
		}},
	}

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			err := Generate(&tc.oto)
			assert.NotNil(t, err)
		})
	}
}

func TestGenerateTempMailError(t *testing.T) {
	opts := options.OreillyTrialOptions{
		CreateUserUrl:        url,
		PasswordRandomLength: 12,
	}

	apiKeyOrig := mailslurp.ApiKey
	mailslurp.ApiKey = "asdasdasdasdasd"

	err := Generate(&opts)
	assert.NotNil(t, err)

	mailslurp.ApiKey = apiKeyOrig
}

// TestGenerateValidArgs function tests if Generate function running properly with proper values
func TestGenerateValidArgs(t *testing.T) {
	cases := []struct {
		caseName string
		oto      options.OreillyTrialOptions
	}{
		{"case1", options.OreillyTrialOptions{
			CreateUserUrl:        url,
			PasswordRandomLength: 12,
		}},
		{"case2", options.OreillyTrialOptions{
			CreateUserUrl:        url,
			PasswordRandomLength: 16,
		}},
	}

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			err := Generate(&tc.oto)
			assert.Nil(t, err)
		})
	}
}
