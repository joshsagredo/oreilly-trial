package oreilly

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/bilalcaliskan/oreilly-trial/internal/logging"
	"github.com/bilalcaliskan/oreilly-trial/internal/options"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	logger *zap.SugaredLogger
	client *http.Client
)

func init() {
	logger = logging.GetLogger()
	client = &http.Client{}
}

// Generate does the heavy lifting, communicates with the Oreilly API
func Generate(opts *options.OreillyTrialOptions, mail, password string) error {
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
	if req, err = http.NewRequest("POST", opts.CreateUserUrl, bytes.NewBuffer(jsonData)); err != nil {
		return errors.Wrap(err, "unable to prepare http request")
	}

	logger.Debug("trying to set request headers")
	setRequestHeaders(req)

	logger.Debug("sending request with http client", "url", opts.CreateUserUrl)
	if resp, err = client.Do(req); err != nil {
		return errors.Wrapf(err, "unable to do http request to remote host %s", opts.CreateUserUrl)
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
	req.Header.Set("authority", "learning.oreilly.com")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("accept", "application/json")
	req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.96 Safari/537.36")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("origin", "https://learning.oreilly.com")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("referer", "https://learning.oreilly.com/p/register/")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
}
