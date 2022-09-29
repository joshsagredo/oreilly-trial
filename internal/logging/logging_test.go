package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetLogger function tests if GetLogger function running properly
func TestGetLogger(t *testing.T) {
	logger := GetLogger()
	assert.NotNil(t, logger)
}

func TestSetLogLevel(t *testing.T) {
	err := SetLogLevel("debug")
	assert.Nil(t, err)
}

func TestSetLogLevelInvalidLevel(t *testing.T) {
	err := SetLogLevel("invalidloglevel")
	assert.NotNil(t, err)
}
