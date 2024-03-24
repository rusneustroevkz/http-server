package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

type Logger interface {
	Logger() *zap.Logger
	Std() *log.Logger

	Info(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
}

type appLog struct {
	log *zap.Logger
	std *log.Logger
	lvl zapcore.Level
}

func NewLogger() Logger {
	l := zap.NewExample()

	return &appLog{
		log: l,
		lvl: zap.DebugLevel,
		std: zap.NewStdLog(l),
	}
}

func (l *appLog) Logger() *zap.Logger {
	return l.log
}

func (l *appLog) Std() *log.Logger {
	return l.std
}
