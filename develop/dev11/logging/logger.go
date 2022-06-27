package logging

import "go.uber.org/zap"

// Logger представляет собой класс-обертку для логгера zap
type Logger struct {
	*zap.SugaredLogger
}

// New возвращает новый Logger
func New() *Logger {
	return &Logger{newZapLogger()}
}

// newZapLogger конфигурирует zap логгер и возвращает его
func newZapLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	return sugar
}
