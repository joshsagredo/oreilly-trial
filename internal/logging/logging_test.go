package logging

import (
	"testing"

	"go.uber.org/zap"
)

func TestGetLogger(t *testing.T) {
	_, err := zap.NewProduction()
	if err != nil {
		t.Errorf("%v\n", err.Error())
	}
}
