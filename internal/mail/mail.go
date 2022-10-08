package mail

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

var url = "https://www.1secmail.com/api/v1/?action=genRandomMailbox"

func GenerateTempMails(count int) ([]string, error) {
	var tempmails []string

	if count > 20 || count < 1 {
		return tempmails, errors.New("invalid attempt count, must be between 1 and 20")
	}

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s&count=%d", url, count), nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return tempmails, errors.Wrap(err, "unable to make request")
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(res.Body)
	body, _ := io.ReadAll(res.Body)

	if res.StatusCode == 200 {
		if err := json.Unmarshal(body, &tempmails); err != nil {
			return tempmails, errors.Wrap(err, "unable to unmarshal json response")
		}
	} else {
		return tempmails, errors.New("API returned non-200 response")
	}

	return tempmails, nil
}
