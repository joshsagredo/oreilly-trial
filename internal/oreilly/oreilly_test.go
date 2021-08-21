package oreilly

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"oreilly-trial/internal/options"
	"oreilly-trial/internal/random"
	"testing"
)

func TestGenerate(t *testing.T) {
	cases := []struct {
		caseName string
		oto      options.OreillyTrialOptions
	}{
		{"case1", options.OreillyTrialOptions{
			CreateUserUrl: "https://learning.oreilly.com/api/v1/registration/individual/",
			EmailDomains:  []string{"jentrix.com"},
			RandomLength:  12,
		}},
		{"case2", options.OreillyTrialOptions{
			CreateUserUrl: "https://learning.oreilly.com/api/v1/registration/individual/",
			EmailDomains:  []string{"geekale.com", "64ge.com"},
			RandomLength:  16,
		}},
	}

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			username := random.GenerateUsername(tc.oto.RandomLength)
			password := random.GeneratePassword(tc.oto.RandomLength)
			emailDomain := random.PickEmail(tc.oto.EmailDomains)
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
				t.Fatalf("%v\n", err.Error())
			}

			req, err := http.NewRequest("POST", tc.oto.CreateUserUrl, bytes.NewBuffer(jsonData))

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

			if resp.StatusCode == 200 {
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
