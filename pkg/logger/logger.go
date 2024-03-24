package logger

import (
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"net/http"
	"time"
)

type Logger interface {
	Logger() *zap.Logger
	Std() *log.Logger
	RequestLogger(next http.Handler) http.Handler

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

func (l *appLog) RequestLogger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()

		defer func() {
			l.Info(
				"request log",
				String("method", r.Method),
				String("path", r.RequestURI),
				Any("request_id", r.Context().Value(middleware.RequestIDKey)),
				Any("duration", time.Since(t1).Milliseconds()),
			)
		}()

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
