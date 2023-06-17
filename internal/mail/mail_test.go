//go:build unit
// +build unit

package mail

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPossiblyValidDomains(t *testing.T) {
	domains, err := GetPossiblyValidDomains()
	assert.NotEmpty(t, domains)
	assert.Nil(t, err)
}

func TestGetPossiblyValidDomainsTimeoutError(t *testing.T) {
	apiURLOrig := ApiURL //nolint:typecheck
	ApiURL = "https://dropmail.p.rapidapi.co/"
	domains, err := GetPossiblyValidDomains()
	assert.Empty(t, domains)
	assert.NotNil(t, err)
	ApiURL = apiURLOrig
}

func TestGetPossiblyValidDomainsUnmarshalError(t *testing.T) {
	apiURLOrig := ApiURL //nolint:typecheck
	response := "non-json response"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprint(w, response); err != nil {
			t.Fatalf("a fatal error occured while writing response body: %s", err.Error())
		}
	}))

	defer func() {
		server.Close()
	}()

	ApiURL = server.URL
	domains, err := GetPossiblyValidDomains()
	assert.Empty(t, domains)
	assert.NotNil(t, err)
	ApiURL = apiURLOrig
}

func TestGetPossiblyValidDomainsEmptyValidDomains(t *testing.T) {
	predefinedValidDomainsOrig := PredefinedValidDomains
	PredefinedValidDomains = []string{}
	domains, err := GetPossiblyValidDomains()
	assert.Empty(t, domains)
	assert.NotNil(t, err)
	PredefinedValidDomains = predefinedValidDomainsOrig
}

func TestGenerateTempMail(t *testing.T) {
	domains, err := GetPossiblyValidDomains()
	assert.NotEmpty(t, domains)
	assert.Nil(t, err)

	mail, err := GenerateTempMail(domains[0])
	assert.NotEmpty(t, mail)
	assert.Nil(t, err)
}

func TestGenerateTempMailTimeoutError(t *testing.T) {
	domains, err := GetPossiblyValidDomains()
	assert.NotEmpty(t, domains)
	assert.Nil(t, err)

	apiURLOrig := ApiURL //nolint:typecheck
	ApiURL = "https://dropmail.p.rapidapi.co/"
	mail, err := GenerateTempMail(domains[0])
	assert.Empty(t, mail)
	assert.NotNil(t, err)
	ApiURL = apiURLOrig
}

func TestGenerateTempMailUnmarshalError(t *testing.T) {
	apiURLOrig := ApiURL //nolint:typecheck
	response := "non-json response"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprint(w, response); err != nil {
			t.Fatalf("a fatal error occured while writing response body: %s", err.Error())
		}
	}))

	defer func() {
		server.Close()
	}()

	ApiURL = server.URL
	mail, err := GenerateTempMail("123456")
	assert.Empty(t, mail)
	assert.NotNil(t, err)
	ApiURL = apiURLOrig
}
