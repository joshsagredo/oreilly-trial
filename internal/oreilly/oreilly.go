package oreilly

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"oreilly-trial/internal/logging"
	"oreilly-trial/internal/options"
	"oreilly-trial/internal/random"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	client *http.Client
)

func init() {
	logger = logging.GetLogger()
	client = &http.Client{}
}

// Generate does the heavy lifting, communicates with the Oreilly API
func Generate(options *options.OreillyTrialOptions) error {
	username := random.GenerateUsername(options.RandomLength)
	password := random.GeneratePassword(options.RandomLength)
	logger.Info("random credentials generated", zap.String("username", username),
		zap.String("password", password))

	emailDomain := random.PickEmail(options.EmailDomains)
	emailAddr := fmt.Sprintf("%s@%s", username, emailDomain)
	values := map[string]string{
		"email":         emailAddr,
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
	jsonData, err := json.Marshal(values)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", options.CreateUserUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	setRequestHeaders(req)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		if err = resp.Body.Close(); err != nil {
			panic(err)
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode == 200 {
		successResponse := successResponse{}
		if err := json.Unmarshal(body, &successResponse); err != nil {
			return err
		}

		logger.Info("trial account successfully created", zap.String("email", emailAddr),
			zap.String("password", password), zap.String("user_id", successResponse.UserID))
	} else {
		logger.Info(string(body))
		return errors.New("an error occurred while creating trial account, please try again")
	}

	return nil
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
