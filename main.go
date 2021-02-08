package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// for more usable domains, check https://temp-mail.org/
	emailDomain := flag.String("emailDomain", "jentrix.com", "usable domain for creating trial " +
		"account, it should be a valid domain")
	length := flag.Int("length", 12, "length of the random generated username and password")
	flag.Parse()

	createUserUrl := "https://learning.oreilly.com/api/v1/user/"
	username := generateUsername(*length)
	password := generatePassword(*length)

	emailAddr := fmt.Sprintf("%s@%s", username, *emailDomain)
	firstName := "John"
	lastName := "Doe"
	values := map[string]string{"email": emailAddr,
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
		log.Panicln(err)
	}

	req, err := http.NewRequest("POST", createUserUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Panicln(err)
	}
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

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()


	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}

	if resp.StatusCode == 201 {
		successResponse := SuccessResponse{}
		err := json.Unmarshal(body, &successResponse)
		if err != nil {
			log.Panicln(err)
		}

		log.Printf("Trial account successfully created!\nemail = %s\npassword = %s\nuser_id = %s\n", emailAddr,
			password, successResponse.UserID)
	} else {
		log.Println("an error occured while creating trial account, please try again!")
		log.Printf("status code = %d\n", resp.StatusCode)
		failureResponse := FailureResponse{}
		err := json.Unmarshal(body, &failureResponse)
		if err != nil {
			log.Panicln(err)
		}

		log.Printf("error messages = %v\n", failureResponse.Email)
	}
}