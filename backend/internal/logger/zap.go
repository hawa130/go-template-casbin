package logger

import "go.uber.org/zap"

var (
	logger *zap.Logger
)

func Init() error {
	l, err := zap.NewProduction()
	if err != nil {
		return err
	}
	logger = l
	return nil
}

func Logger() *zap.Logger {
	return logger
}

func Sync() {
	if logger != nil {
		_ = logger.Sync()
	}
}
