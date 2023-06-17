//go:build unit
// +build unit

package oreilly

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bilalcaliskan/oreilly-trial/internal/logging"
	"github.com/bilalcaliskan/oreilly-trial/internal/mail"
	"github.com/bilalcaliskan/oreilly-trial/internal/random"

	"github.com/stretchr/testify/assert"
)

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

	apiURLOrig := apiURL
	apiURL = server.URL
	err := Generate("notreallyrequiredmail@example.com", "123123123123", logging.GetLogger())
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), expectedError)

	apiURL = apiURLOrig
}

// TestGenerateInvalidHost function tests if Generate function fails on broken Host argument
func TestGenerateInvalidHost(t *testing.T) {
	expectedError := "no such host"

	apiURLOrig := apiURL
	apiURL = "https://foo.example.com/"

	err := Generate("notreallyrequiredmail@example.com", "123123123123", logging.GetLogger())
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), expectedError)

	apiURL = apiURLOrig
}

// TestGenerateValidArgs function tests if Generate function running properly with proper values
func TestGenerateValidArgs(t *testing.T) {
	password, err := random.GeneratePassword()
	assert.NotEmpty(t, password)
	assert.Nil(t, err)

	domains, _ := mail.GetPossiblyValidDomains()

	for _, id := range domains {
		email, err := mail.GenerateTempMail(id)
		assert.NotEmpty(t, email)
		assert.Nil(t, err)

		err = Generate(email, password, logging.GetLogger())

		if err == nil {
			break
		}
	}

	assert.Nil(t, err)
}
