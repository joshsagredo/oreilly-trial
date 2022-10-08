package oreilly

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bilalcaliskan/oreilly-trial/internal/logging"
	"github.com/bilalcaliskan/oreilly-trial/internal/mail"
	"github.com/bilalcaliskan/oreilly-trial/internal/random"
	"go.uber.org/zap"

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

	err := Generate(&oto, "notreallyrequiredmail@example.com", "123123123123")
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
		AttemptCount:         10,
	}

	err := Generate(&oto, "notreallyrequiredmail@example.com", "123123123123")
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
			AttemptCount:         10,
		}},
		{"case2", options.OreillyTrialOptions{
			CreateUserUrl:        url,
			PasswordRandomLength: 665,
			AttemptCount:         10,
		}},
	}

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			err := Generate(&tc.oto, "notreallyrequiredmail@example.com", "123123123123")
			assert.NotNil(t, err)
			assert.NotNil(t, err)
		})
	}
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
			AttemptCount:         10,
		}},
	}

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			password, err := random.GeneratePassword(tc.oto.PasswordRandomLength)
			assert.NotEmpty(t, password)
			assert.Nil(t, err)

			tempmails, err := mail.GenerateTempMails(tc.oto.AttemptCount)
			if err != nil {
				logging.GetLogger().Error("an error occurred while generating temp mails", zap.String("error", err.Error()))
				return
			}

			for _, temp := range tempmails {
				err = Generate(&tc.oto, temp, password)

				if err == nil {
					break
				}
			}

			assert.Nil(t, err)
		})
	}
}
