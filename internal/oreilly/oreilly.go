package oreilly

import (
	"bytes"
	"encoding/json"
	"github.com/rs/zerolog"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

var (
	client *http.Client
	apiURL = "https://learning.oreilly.com/api/v1/registration/individual/"
)

func init() {
	client = &http.Client{}
}

// Generate does the heavy lifting, communicates with the Oreilly API
func Generate(mail, password string, logger zerolog.Logger) error {
	var (
		jsonData []byte
		req      *http.Request
		resp     *http.Response
		respBody []byte
		err      error
	)

	// prepare json data
	values := map[string]string{
		"email":         mail,
		"password":      password,
		"first_name":    "John",
		"last_name":     "Doe",
		"country":       "US",
		"t_c_agreement": "true",
		"contact":       "true",
		"trial_length":  "10",
		"path":          "/register/",
		"source":        "payments-client-register",
	}

	// marshall the json body
	if jsonData, err = json.Marshal(values); err != nil {
		return errors.Wrap(err, "unable to marshal request body")
	}

	// prepare and make the request
	if req, err = http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData)); err != nil {
		return errors.Wrap(err, "unable to prepare http request")
	}

	logger.Debug().Msg("trying to set request headers")
	setRequestHeaders(req)

	logger.Debug().Str("url", apiURL).Msg("sending request with http client")
	if resp, err = client.Do(req); err != nil {
		return errors.Wrapf(err, "unable to do http request to remote host %s", apiURL)
	}

	defer func(body io.ReadCloser) {
		err = body.Close()
	}(resp.Body)

	// read the response
	if respBody, err = io.ReadAll(resp.Body); err != nil {
		return errors.Wrap(err, "unable to read response")
	}

	if resp.StatusCode == 200 {
		var successResponse successResponse
		if err = json.Unmarshal(respBody, &successResponse); err != nil {
			return errors.Wrap(err, "unable to unmarshal json response")
		}
	} else {
		return errors.New(string(respBody))
	}

	return err
}

// setRequestHeaders gets the http.Request as input and add some headers for proper API request
func setRequestHeaders(req *http.Request) {
	req.Header.Set("authority", requestAuthority)
	req.Header.Set("pragma", requestPragma)
	req.Header.Set("cache-control", requestCacheControl)
	req.Header.Set("accept", requestAccept)
	req.Header.Set("user-agent", requestUserAgent)
	req.Header.Set("content-type", requestContentType)
	req.Header.Set("origin", requestOrigin)
	req.Header.Set("sec-fetch-site", requestSecFetchSite)
	req.Header.Set("sec-fetch-mode", requestSecFetchMode)
	req.Header.Set("sec-fetch-dest", requestSecFetchDest)
	req.Header.Set("referer", requestReferer)
	req.Header.Set("accept-language", requestAcceptLang)
}
