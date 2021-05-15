package oreilly

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"oreilly-trial/pkg/random"
	"testing"
)

func TestGenerate(t *testing.T) {
	cases := []struct {
		caseName, username, password string
	}{
		{"case1", random.GenerateUsername(randomLength),
			random.GeneratePassword(randomLength)},
		{"case2", random.GenerateUsername(randomLength),
			random.GeneratePassword(randomLength)},
	}

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			emailAddr := fmt.Sprintf("%s@%s", tc.username, emailDomain)
			values := map[string]string{
				"email":         emailAddr,
				"password":      tc.password,
				"first_name":    "John",
				"last_name":     "Doe",
				"country":       "US",
				"t_c_agreement": "true",
				"contact":       "true",
				"path":          "/register/",
				"source":        "payments-client-register",
			}

			jsonData, err := json.Marshal(values)
			if err != nil {
				t.Fatalf("%v\n", err.Error())
			}

			req, err := http.NewRequest("POST", createUserUrl, bytes.NewBuffer(jsonData))
			if err != nil {
				t.Fatalf("%v\n", err.Error())
			}

			setRequestHeaders(req)

			resp, err := client.Do(req)
			if err != nil {
				t.Fatalf("%v\n", err.Error())
			}

			defer func() {
				err = resp.Body.Close()
				if err != nil {
					t.Fatalf("%v\n", err.Error())
				}
			}()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("%v\n", err.Error())
			}

			if resp.StatusCode == 201 {
				successResponse := successResponse{}
				err := json.Unmarshal(body, &successResponse)
				if err != nil {
					t.Fatalf("%v\n", err.Error())
				}

				t.Logf("trial account successfully created")
			} else {
				t.Fatalf("an error occurred while creating user")
			}
		})
	}
}
