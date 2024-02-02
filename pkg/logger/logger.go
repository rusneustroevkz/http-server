package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Logger() *zap.Logger

	Info(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
}

type log struct {
	log *zap.Logger
	lvl zapcore.Level
}

func NewLogger() Logger {
	return &log{
		log: zap.NewExample(),
		lvl: zap.DebugLevel,
	}
}

func (l *log) Logger() *zap.Logger {
	return l.log
}
