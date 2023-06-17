//go:build unit
// +build unit

package prompt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPromptRunner(t *testing.T) {
	runner := GetPromptRunner()
	assert.NotNil(t, runner)
}

func TestGetSelectRunner(t *testing.T) {
	runner := GetSelectRunner()
	assert.NotNil(t, runner)
}
