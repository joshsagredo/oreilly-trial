package mailslurp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateTempMail(t *testing.T) {
	email, err := GenerateTempMail()
	assert.NotNil(t, email)
	assert.NotEmpty(t, email)
	assert.Nil(t, err)
}
