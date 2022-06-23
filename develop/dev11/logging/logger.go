package logging

import "go.uber.org/zap"

type Logger struct {
	*zap.SugaredLogger
}

func New() *Logger {
	return &Logger{newZapLogger()}
}

func newZapLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	return sugar
}
