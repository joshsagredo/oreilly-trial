package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
	Atomic zap.AtomicLevel
)

func init() {
	Atomic = zap.NewAtomicLevel()
	Atomic.SetLevel(zap.InfoLevel)
	logger = zap.New(zapcore.NewTee(zapcore.NewCore(zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:   "message",
		LevelKey:     "severity",
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		TimeKey:      "time",
		EncodeTime:   zapcore.RFC3339TimeEncoder,
		CallerKey:    "caller",
		EncodeCaller: zapcore.FullCallerEncoder,
	}), zapcore.Lock(os.Stdout), Atomic)))
}

// GetLogger returns the shared *zap.Logger
func GetLogger() *zap.Logger {
	return logger
}
