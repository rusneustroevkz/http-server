package logger

import "go.uber.org/zap"

type Field zap.Field

func String(key, msg string) Field {
	return Field(zap.String(key, msg))
}

func Error(err error) Field {
	return Field(zap.Error(err))
}

func Int64(key string, value int64) Field {
	return Field(zap.Int64(key, value))
}

func Any(key string, value any) Field {
	return Field(zap.Any(key, value))
}
