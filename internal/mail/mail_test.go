package mail

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPossiblyValidDomains(t *testing.T) {
	urlOrig := url
	// url = "https://www.1secmail.com/api/v1/?action=genRandomMailboxxxxxx"
	mails, err := GetPossiblyValidDomains()
	assert.NotEmpty(t, mails)
	assert.Nil(t, err)
	url = urlOrig
}

func TestGenerateTempMail(t *testing.T) {

}
