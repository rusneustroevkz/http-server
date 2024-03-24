package logger

import (
	"net/http"
	"time"
)

type entryFields struct {
	path      string
	requestID any
	method    string
}

type logEntry struct {
	log Logger
}

func (l *logEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	switch e := extra.(type) {
	case entryFields:
		l.log.Info(
			"request log",
			String("method", e.method),
			String("path", e.path),
			Int("status", status),
			Int64("duration", elapsed.Milliseconds()),
			Any("request_id", e.requestID),
		)
	default:
		l.log.Error("cannot log request, undefined type")
	}
}

func (l *logEntry) Panic(v interface{}, stack []byte) {
	l.log.Error("has panic", Any("message", v), String("stack", string(stack)))
}
