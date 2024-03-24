package logger

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Logger() *zap.Logger
	Std() *log.Logger
	NewLogEntry(r *http.Request) middleware.LogEntry
	RequestLogger(f middleware.LogFormatter) func(next http.Handler) http.Handler

	Info(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
	Error(msg string, fields ...Field)
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

func (l *appLog) NewLogEntry(r *http.Request) middleware.LogEntry {
	return &logEntry{
		log: l,
	}
}

func (l *appLog) RequestLogger(f middleware.LogFormatter) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			entry := f.NewLogEntry(r)
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			defer func() {
				entry.Write(ww.Status(), ww.BytesWritten(), ww.Header(), time.Since(t1), entryFields{
					path:      r.RequestURI,
					requestID: r.Context().Value(middleware.RequestIDKey),
					method:    r.Method,
				})
			}()

			next.ServeHTTP(ww, middleware.WithLogEntry(r, entry))
		}
		return http.HandlerFunc(fn)
	}
}
