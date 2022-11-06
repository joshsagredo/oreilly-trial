package mail

import (
	"encoding/json"
	"fmt"
	"github.com/bilalcaliskan/oreilly-trial/internal/logging"
	"io"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

var (
	url   = "https://dropmail.p.rapidapi.com/"
	token = "none"
)

func GetPossiblyValidDomains() ([]string, error) {
	var possibleValidDomains []string
	var domainRequestData = strings.NewReader("{\"query\":\"query domains { domains { id name introducedAt availableVia }}\",\"variables\":{}}")
	logging.GetLogger().Infow("token print", "token", token)
	domainRequest, _ := http.NewRequest("POST", url, domainRequestData)
	domainRequest.Header.Add("content-type", "application/json")
	domainRequest.Header.Add("X-RapidAPI-Key", token)
	domainRequest.Header.Add("X-RapidAPI-Host", "dropmail.p.rapidapi.com")
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
		if v.Name == "mimimail.me" {
			possibleValidDomains = append(possibleValidDomains, v.ID)
		}
		// TODO: find more valid domains
	}

	if len(possibleValidDomains) == 0 {
		return possibleValidDomains, errors.New("empty list of valid domains")
	}

	return possibleValidDomains, nil
}

func GenerateTempMail(domainID string) (string, error) {
	var emailRequestData = strings.NewReader(fmt.Sprintf("{\"query\":\"mutation introduceSession($input: IntroduceSessionInput) { "+
		"introduceSession(input: $input) { "+
		"id "+
		"addresses { "+
		"address "+
		"} "+
		"expiresAt "+
		"} "+
		"}\",\"variables\":{\"input\":{\"withAddress\":true,\"domainId\":\"%s\"}}}", domainID))
	emailRequest, _ := http.NewRequest("POST", url, emailRequestData)
	emailRequest.Header.Add("content-type", "application/json")
	emailRequest.Header.Add("X-RapidAPI-Key", "6a3f9418famshdeeac0fe2f34a7cp1fd1c1jsn26245029950d")
	emailRequest.Header.Add("X-RapidAPI-Host", "dropmail.p.rapidapi.com")
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

	return resp.Addresses[0].Address, nil
}
