package mail

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var (
	ApiURL = "https://dropmail.p.rapidapi.com/"
	//token                  = "none"
	token                  = "6a3f9418famshdeeac0fe2f34a7cp1fd1c1jsn26245029950d"
	PredefinedValidDomains = []string{"mailpwr.com", "mimimail.me"}
)

func GetPossiblyValidDomains() ([]string, error) {
	var possibleValidDomains []string
	var domainRequestData = strings.NewReader(domainRequestQuery)
	domainRequest, _ := http.NewRequest("POST", ApiURL, domainRequestData)
	domainRequest.Header.Add("content-type", contentType)
	domainRequest.Header.Add("X-RapidAPI-Key", token)
	domainRequest.Header.Add("X-RapidAPI-Host", hostHeader)
	domainResp, err := http.DefaultClient.Do(domainRequest)
	if err != nil {
		return possibleValidDomains, err
	}

	defer func() {
		if err := domainResp.Body.Close(); err != nil {
			panic(err)
		}
	}()
	body, err := io.ReadAll(domainResp.Body)
	if err != nil {
		return possibleValidDomains, err
	}

	var domainResponse DomainResponse
	if err := json.Unmarshal(body, &domainResponse); err != nil {
		return possibleValidDomains, err
	}

	for _, v := range domainResponse.Domains {
		if contains(PredefinedValidDomains, v.Name) {
			possibleValidDomains = append(possibleValidDomains, v.ID)
		}
	}

	if len(possibleValidDomains) == 0 {
		return possibleValidDomains, errors.New("empty list of valid domains")
	}

	return possibleValidDomains, nil
}

func GenerateTempMail(domainID string) (string, error) {
	var emailRequestData = strings.NewReader(fmt.Sprintf(emailRequestQuery, domainID))
	emailRequest, _ := http.NewRequest("POST", ApiURL, emailRequestData)
	emailRequest.Header.Add("content-type", contentType)
	emailRequest.Header.Add("X-RapidAPI-Key", token)
	emailRequest.Header.Add("X-RapidAPI-Host", hostHeader)
	res, err := http.DefaultClient.Do(emailRequest)
	if err != nil {
		return "", err
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			panic(err)
		}
	}()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var resp EmailResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return "", err
	}

	if len(resp.Addresses) == 0 {
		return "", errors.New("no email returned from API")
	}

	return resp.Addresses[0].Address, nil
}
