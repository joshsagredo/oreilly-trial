//go:build unit
// +build unit

package generator

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/bilalcaliskan/oreilly-trial/internal/mail"
	"github.com/stretchr/testify/assert"
)

func TestRunGenerator(t *testing.T) {
	err := RunGenerator()
	assert.Nil(t, err)
}

func TestRunGeneratorDomainError(t *testing.T) {
	apiURLOrig := mail.ApiURL //nolint:typecheck
	mail.ApiURL = "https://dropmail.p.rapidapi.co/"

	err := RunGenerator()
	assert.NotNil(t, err)

	mail.ApiURL = apiURLOrig
}

func TestRunGeneratorGenerateTempMailError(t *testing.T) {
	apiURLOrig := mail.ApiURL //nolint:typecheck
	response := "{\"data\":{\"domains\":[{\"name\":\"10mail.org\",\"introducedAt\":\"2013-11-13T11:00:00.000+00:00\"," +
		"\"id\":\"RG9tYWluOjI\",\"availableVia\":[\"APP\",\"API\",\"TELEGRAM\",\"VIBER\",\"WEB\"]},{\"name\":\"10mail.tk\"" +
		",\"introducedAt\":\"2021-01-12T11:00:00.000+00:00\",\"id\":\"RG9tYWluOjE2\",\"availableVia\":[\"APP\",\"API\"" +
		",\"TELEGRAM\",\"VIBER\",\"WEB\"]},{\"name\":\"dropmail.me\",\"introducedAt\":\"2013-05-10T10:00:00.000+00:00\"" +
		",\"id\":\"RG9tYWluOjE\",\"availableVia\":[\"APP\",\"API\",\"TELEGRAM\",\"VIBER\",\"WEB\"]},{\"name\":\"emlhub.com\"" +
		",\"introducedAt\":\"2017-05-14T10:00:00.000+00:00\",\"id\":\"RG9tYWluOjg\",\"availableVia\":[\"APP\",\"API\",\"TELEGRAM\"" +
		",\"VIBER\",\"WEB\"]},{\"name\":\"emlpro.com\",\"introducedAt\":\"2017-05-14T10:00:00.000+00:00\",\"id\":\"RG9tYWluOjc\"" +
		",\"availableVia\":[\"APP\",\"API\",\"TELEGRAM\",\"VIBER\",\"WEB\"]},{\"name\":\"emltmp.com\",\"introducedAt\"" +
		":\"2016-05-20T10:00:00.000+00:00\",\"id\":\"RG9tYWluOjY\",\"availableVia\":[\"APP\",\"API\",\"TELEGRAM\",\"VIBER\"" +
		",\"WEB\"]},{\"name\":\"firste.ml\",\"introducedAt\":\"2019-10-02T10:00:00.000+00:00\",\"id\":\"RG9tYWluOjEy\"" +
		",\"availableVia\":[\"APP\",\"API\",\"TELEGRAM\",\"VIBER\",\"WEB\"]},{\"name\":\"flymail.tk\",\"introducedAt\":\"2021-09-16T21:24:30.019+00:00\"" +
		",\"id\":\"RG9tYWluOjE4\",\"availableVia\":[\"APP\",\"API\",\"TELEGRAM\",\"VIBER\",\"WEB\"]},{\"name\":\"freeml.net\"" +
		",\"introducedAt\":\"2021-01-12T11:00:00.000+00:00\",\"id\":\"RG9tYWluOjE1\",\"availableVia\":[\"APP\",\"API\",\"TELEGRAM\"" +
		",\"VIBER\",\"WEB\"]},{\"name\":\"laste.ml\",\"introducedAt\":\"2019-10-02T10:00:00.000+00:00\",\"id\":\"RG9tYWluOjEz\"" +
		",\"availableVia\":[\"APP\",\"API\",\"TELEGRAM\",\"VIBER\",\"WEB\"]},{\"name\":\"mailpwr.com\",\"introducedAt\":\"2021-10-06T22:18:11.552+00:00\"" +
		",\"id\":\"RG9tYWluOjIw\",\"availableVia\":[\"APP\",\"API\"]},{\"name\":\"mimimail.me\",\"introducedAt\":\"2022-01-11T01:52:41.239+00:00\"" +
		",\"id\":\"RG9tYWluOjIx\",\"availableVia\":[\"APP\",\"API\"]},{\"name\":\"minimail.gq\",\"introducedAt\":\"2021-09-16T21:22:16.896+00:00\"" +
		",\"id\":\"RG9tYWluOjE3\",\"availableVia\":[\"APP\",\"API\",\"TELEGRAM\",\"VIBER\",\"WEB\"]},{\"name\":\"spymail.one\"" +
		",\"introducedAt\":\"2021-10-05T11:05:59.918+00:00\",\"id\":\"RG9tYWluOjE5\",\"availableVia\":[\"APP\",\"API\",\"TELEGRAM\"" +
		",\"VIBER\",\"WEB\"]},{\"name\":\"yomail.info\",\"introducedAt\":\"2014-10-30T11:00:00.000+00:00\",\"id\":\"RG9tYWluOjQ\"" +
		",\"availableVia\":[\"APP\",\"API\",\"TELEGRAM\",\"VIBER\",\"WEB\"]},{\"name\":\"zeroe.ml\",\"introducedAt\":\"2019-10-02T10:00:00.000+00:00\"" +
		",\"id\":\"RG9tYWluOjEx\",\"availableVia\":[\"APP\",\"API\",\"TELEGRAM\",\"VIBER\",\"WEB\"]}]}}"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprint(w, response); err != nil {
			t.Fatalf("a fatal error occured while writing response body: %s", err.Error())
		}
	}))

	defer func() {
		server.Close()
	}()

	mail.ApiURL = server.URL
	err := RunGenerator()
	assert.NotNil(t, err)

	mail.ApiURL = apiURLOrig
}

func TestRunGeneratorTrialAccountCreateError(t *testing.T) {
	apiURLOrig := mail.ApiURL //nolint:typecheck

	rr := newResponseWriter()
	server := httptest.NewServer(handlerResponse(rr))
	defer func() {
		server.Close()
	}()

	mail.ApiURL = server.URL
	err := RunGenerator()
	assert.NotNil(t, err)

	mail.ApiURL = apiURLOrig
}

type responseWriter struct {
	resp  map[int]string
	count int
	lock  *sync.Mutex
}

func newResponseWriter() *responseWriter {
	r := new(responseWriter)
	r.lock = new(sync.Mutex)
	r.resp = map[int]string{
		0: "{\"data\":{\"introduceSession\":{\"id\":\"U2Vzc2lvbjqlmFu26iZORqMlWdOy3DCC\",\"expiresAt\":\"2022-11-12T13:19:17+00:00\",\"addresses\":[{\"address\":\"asdasfasfas@mailpwr.com\"}]}}}",
		1: "{\"data\":{\"domains\":[{\"name\":\"10mail.org\",\"introducedAt\":\"2013-11-13T11:00:00.000+00:00\"," +
			"\"id\":\"RG9tYWluOjI\",\"availableVia\":[\"APP\",\"API\",\"TELEGRAM\",\"VIBER\",\"WEB\"]},{\"name\":\"10mail.tk\"" +
			",\"introducedAt\":\"2021-01-12T11:00:00.000+00:00\",\"id\":\"RG9tYWluOjE2\",\"availableVia\":[\"APP\",\"API\"" +
			",\"TELEGRAM\",\"VIBER\",\"WEB\"]},{\"name\":\"dropmail.me\",\"introducedAt\":\"2013-05-10T10:00:00.000+00:00\"" +
			",\"id\":\"RG9tYWluOjE\",\"availableVia\":[\"APP\",\"API\",\"TELEGRAM\",\"VIBER\",\"WEB\"]},{\"name\":\"emlhub.com\"" +
			",\"introducedAt\":\"2017-05-14T10:00:00.000+00:00\",\"id\":\"RG9tYWluOjg\",\"availableVia\":[\"APP\",\"API\",\"TELEGRAM\"" +
			",\"VIBER\",\"WEB\"]},{\"name\":\"emlpro.com\",\"introducedAt\":\"2017-05-14T10:00:00.000+00:00\",\"id\":\"RG9tYWluOjc\"" +
			",\"availableVia\":[\"APP\",\"API\",\"TELEGRAM\",\"VIBER\",\"WEB\"]},{\"name\":\"emltmp.com\",\"introducedAt\"" +
			":\"2016-05-20T10:00:00.000+00:00\",\"id\":\"RG9tYWluOjY\",\"availableVia\":[\"APP\",\"API\",\"TELEGRAM\",\"VIBER\"" +
			",\"WEB\"]},{\"name\":\"firste.ml\",\"introducedAt\":\"2019-10-02T10:00:00.000+00:00\",\"id\":\"RG9tYWluOjEy\"" +
			",\"availableVia\":[\"APP\",\"API\",\"TELEGRAM\",\"VIBER\",\"WEB\"]},{\"name\":\"flymail.tk\",\"introducedAt\":\"2021-09-16T21:24:30.019+00:00\"" +
			",\"id\":\"RG9tYWluOjE4\",\"availableVia\":[\"APP\",\"API\",\"TELEGRAM\",\"VIBER\",\"WEB\"]},{\"name\":\"freeml.net\"" +
			",\"introducedAt\":\"2021-01-12T11:00:00.000+00:00\",\"id\":\"RG9tYWluOjE1\",\"availableVia\":[\"APP\",\"API\",\"TELEGRAM\"" +
			",\"VIBER\",\"WEB\"]},{\"name\":\"laste.ml\",\"introducedAt\":\"2019-10-02T10:00:00.000+00:00\",\"id\":\"RG9tYWluOjEz\"" +
			",\"availableVia\":[\"APP\",\"API\",\"TELEGRAM\",\"VIBER\",\"WEB\"]},{\"name\":\"mailpwr.com\",\"introducedAt\":\"2021-10-06T22:18:11.552+00:00\"" +
			",\"id\":\"RG9tYWluOjIw\",\"availableVia\":[\"APP\",\"API\"]},{\"name\":\"mimimail.me\",\"introducedAt\":\"2022-01-11T01:52:41.239+00:00\"" +
			",\"id\":\"RG9tYWluOjIx\",\"availableVia\":[\"APP\",\"API\"]},{\"name\":\"minimail.gq\",\"introducedAt\":\"2021-09-16T21:22:16.896+00:00\"" +
			",\"id\":\"RG9tYWluOjE3\",\"availableVia\":[\"APP\",\"API\",\"TELEGRAM\",\"VIBER\",\"WEB\"]},{\"name\":\"spymail.one\"" +
			",\"introducedAt\":\"2021-10-05T11:05:59.918+00:00\",\"id\":\"RG9tYWluOjE5\",\"availableVia\":[\"APP\",\"API\",\"TELEGRAM\"" +
			",\"VIBER\",\"WEB\"]},{\"name\":\"yomail.info\",\"introducedAt\":\"2014-10-30T11:00:00.000+00:00\",\"id\":\"RG9tYWluOjQ\"" +
			",\"availableVia\":[\"APP\",\"API\",\"TELEGRAM\",\"VIBER\",\"WEB\"]},{\"name\":\"zeroe.ml\",\"introducedAt\":\"2019-10-02T10:00:00.000+00:00\"" +
			",\"id\":\"RG9tYWluOjEx\",\"availableVia\":[\"APP\",\"API\",\"TELEGRAM\",\"VIBER\",\"WEB\"]}]}}",
		2: "{\"data\":{\"introduceSession\":{\"id\":\"U2Vzc2lvbjqlmFu26iZORqMlWdOy3DCC\",\"expiresAt\":\"2022-11-12T13:19:17+00:00\",\"addresses\":[{\"address\":\"asdasfas@mailpwr.com\"}]}}}",
	}
	r.count = 0
	return r
}

func (r *responseWriter) getResp() string {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.count++
	return r.resp[r.count%3]
}

func handlerResponse(rr *responseWriter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(rr.getResp()))
	})
}
