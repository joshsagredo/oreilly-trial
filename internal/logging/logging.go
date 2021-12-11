package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var logger *zap.Logger

func init() {
	logger = zap.New(zapcore.NewTee(zapcore.NewCore(zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:   "message",
		LevelKey:     "severity",
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		TimeKey:      "time",
		EncodeTime:   zapcore.RFC3339TimeEncoder,
		CallerKey:    "caller",
		EncodeCaller: zapcore.FullCallerEncoder,
	}), zapcore.Lock(os.Stdout), zap.InfoLevel)))
}

// GetLogger returns the shared *zap.Logger
func GetLogger() *zap.Logger {
	return logger
}
