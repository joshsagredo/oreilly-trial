package oreilly

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	flag "github.com/spf13/pflag"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"oreilly-trial/pkg/random"
)

var (
	logger *zap.Logger
	client *http.Client
	err error
	createUserUrl, emailDomain string
	randomLength int
)

func init() {
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}

	client = &http.Client{}
	flag.StringVar(&createUserUrl, "createUserUrl", "https://learning.oreilly.com/api/v1/user/",
		"url of the user creation on Oreilly API")
	// for more usable domains, check https://temp-mail.org/
	flag.StringVar(&emailDomain, "emailDomain", "jentrix.com", "usable domain for creating trial " +
		"account, it should be a valid domain")
	flag.IntVar(&randomLength, "length", 12, "length of the random generated username and password")
	flag.Parse()
}

func Generate() error {
	username := random.GenerateUsername(randomLength)
	password := random.GeneratePassword(randomLength)
	logger.Info("random credentials generated", zap.String("username", username),
		zap.String("password", password))

	emailAddr := fmt.Sprintf("%s@%s", username, emailDomain)
	values := map[string]string{
		"email": emailAddr,
		"password": password,
		"first_name": "John",
		"last_name": "Doe",
		"country": "US",
		"t_c_agreement": "true",
		"contact": "true",
		"path": "/register/",
		"source": "payments-client-register",
	}
	jsonData, err := json.Marshal(values)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", createUserUrl, bytes.NewBuffer(jsonData))
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

	if resp.StatusCode == 201 {
		successResponse := successResponse{}
		err := json.Unmarshal(body, &successResponse)
		if err != nil {
			return err
		}

		logger.Info("trial account successfully created", zap.String("email", emailAddr),
			zap.String("password", password), zap.String("user_id", successResponse.UserID))
	} else {
		return errors.New("an error occured while creating trial account, please try again")
	}
	return nil
}