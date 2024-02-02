package logger

import "go.uber.org/zap"

type Field zap.Field

func String(key, msg string) Field {
	return Field(zap.String(key, msg))
}

func Error(err error) Field {
	return Field(zap.Error(err))
}
