package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	flag "github.com/spf13/pflag"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	http2 "oreilly-trial/pkg/http"
	"oreilly-trial/pkg/random"
)

var (
	logger *zap.Logger
	err error
	emailDomain string
	length int
)

func init() {
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}

	// for more usable domains, check https://temp-mail.org/
	flag.StringVar(&emailDomain, "emailDomain", "jentrix.com", "usable domain for creating trial " +
		"account, it should be a valid domain")
	flag.IntVar(&length, "length", 12, "length of the random generated username and password")
	flag.Parse()
}

func main() {
	createUserUrl := "https://learning.oreilly.com/api/v1/user/"
	username := random.GenerateUsername(length)
	password := random.GeneratePassword(length)
	logger.Info("random credentials generated", zap.String("username", username),
		zap.String("password", password))

	emailAddr := fmt.Sprintf("%s@%s", username, emailDomain)
	firstName := "John"
	lastName := "Doe"
	values := map[string]string{
		"email": emailAddr,
		"password": password,
		"first_name": firstName,
		"last_name": lastName,
		"country": "US",
		"t_c_agreement": "true",
		"contact": "true",
		"path": "/register/",
		"source": "payments-client-register",
	}
	jsonData, err := json.Marshal(values)
	if err != nil {
		logger.Fatal("fatal error occured while marshaling request body", zap.String("error", err.Error()))
	}

	req, err := http.NewRequest("POST", createUserUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		logger.Fatal("fatal error occured while creating new POST request", zap.String("error", err.Error()))
	}

	http2.SetRequestHeaders(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Fatal("fatal error occured while making HTTP request", zap.String("error", err.Error()))
	}

	defer func() {
		err = resp.Body.Close()
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Fatal("fatal error occured while reading response body", zap.String("error", err.Error()))
	}

	if resp.StatusCode == 201 {
		successResponse := http2.SuccessResponse{}
		err := json.Unmarshal(body, &successResponse)
		if err != nil {
			logger.Fatal("fatal error occured while unmarshaling response body", zap.String("error", err.Error()))
		}

		logger.Info("trial account successfully created", zap.String("email", emailAddr),
			zap.String("password", password), zap.String("user_id", successResponse.UserID))
	} else {
		logger.Error("an error occured while creating trial account, please try again!")
		failureResponse := http2.FailureResponse{}
		err := json.Unmarshal(body, &failureResponse)
		if err != nil {
			logger.Fatal("fatal error occured while unmarshaling response body", zap.String("error", err.Error()))
		}
	}
}