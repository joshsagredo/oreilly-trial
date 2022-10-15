package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.SugaredLogger
	atomic zap.AtomicLevel
)

func init() {
	atomic = zap.NewAtomicLevel()
	atomic.SetLevel(zap.InfoLevel)
	logger = zap.New(zapcore.NewTee(zapcore.NewCore(zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:   "message",
		LevelKey:     "severity",
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		TimeKey:      "time",
		EncodeTime:   zapcore.RFC3339TimeEncoder,
		CallerKey:    "caller",
		EncodeCaller: zapcore.FullCallerEncoder,
	}), zapcore.Lock(os.Stdout), atomic))).Sugar()
}

// GetLogger returns the shared *zap.Logger
func GetLogger() *zap.SugaredLogger {
	return logger
}

func SetLogLevel(level string) error {
	parsedLevel, err := zapcore.ParseLevel(level)
	if err != nil {
		return err
	}

	atomic.SetLevel(parsedLevel)
	return nil
}
