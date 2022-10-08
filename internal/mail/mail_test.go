package mail

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateTempMailInvalidUrl(t *testing.T) {
	urlOrig := url
	url = "https://www.1secmail.com/api/v1/?action=genRandomMailboxxxxxx"
	mails, err := GenerateTempMails(1)
	assert.Nil(t, mails)
	assert.NotNil(t, err)
	url = urlOrig
}

func TestGenerateTempMail(t *testing.T) {
	mails, err := GenerateTempMails(1)
	assert.NotNil(t, mails)
	assert.Nil(t, err)
}

func TestGenerateTempMailInvalidAttemptCount(t *testing.T) {
	mails, err := GenerateTempMails(100)
	assert.Nil(t, mails)
	assert.NotNil(t, err)
}
