package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var logger *zap.Logger

func init() {
	cfgConsole := zapcore.EncoderConfig{
		MessageKey:   "message",
		LevelKey:     "severity",
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		TimeKey:      "time",
		EncodeTime:   zapcore.RFC3339TimeEncoder,
		CallerKey:    "caller",
		EncodeCaller: zapcore.FullCallerEncoder,
	}

	core := zapcore.NewTee(zapcore.NewCore(zapcore.NewJSONEncoder(cfgConsole), zapcore.Lock(os.Stdout), zap.InfoLevel))
	logger = zap.New(core)
	logger.Info("An info level message")
}

// GetLogger returns the shared *zap.Logger
func GetLogger() *zap.Logger {
	return logger
}
