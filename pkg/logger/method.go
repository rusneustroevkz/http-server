package logger

import "go.uber.org/zap"

func (l *log) Info(msg string, fields ...Field) {
	zf := make([]zap.Field, 0, cap(fields))
	for _, field := range fields {
		zf = append(zf, zap.Field(field))
	}
	l.log.Info(msg, zf...)
}

func (l *log) Fatal(msg string, fields ...Field) {
	zf := make([]zap.Field, 0, cap(fields))
	for _, field := range fields {
		zf = append(zf, zap.Field(field))
	}
	l.log.Fatal(msg, zf...)
}
