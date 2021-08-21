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
	logger.Info(emailDomain)
	emailAddr := fmt.Sprintf("%s@%s", username, emailDomain)
	logger.Info(emailAddr)
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
		err = resp.Body.Close()
		if err != nil {
			panic(err)
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode == 200 {
		successResponse := successResponse{}
		err := json.Unmarshal(body, &successResponse)
		if err != nil {
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
