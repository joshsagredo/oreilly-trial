package logging

import (
	"testing"
)

func TestGetLogger(t *testing.T) {
	t.Log("getting logger")
	logger := GetLogger()
	t.Log("will try logger for debugging")
	logger.Info("this is a test log by *zap.Logger!")
}
