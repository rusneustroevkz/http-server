package metrics

import (
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"net/http"
	"time"
)

const (
	namespace = "test"
	subsystem = "backend_example"
)

var (
	httpRequestHistogram = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace:                       namespace,
			Subsystem:                       subsystem,
			Name:                            "http_request_histogram",
			Help:                            "",
			ConstLabels:                     nil,
			Buckets:                         prometheus.DefBuckets,
			NativeHistogramBucketFactor:     0,
			NativeHistogramZeroThreshold:    0,
			NativeHistogramMaxBucketNumber:  0,
			NativeHistogramMinResetDuration: 0,
			NativeHistogramMaxZeroThreshold: 0,
		},
		[]string{"method", "path", "status"},
	)
)

func httpRequestCountInc(method, path string, status int, start time.Time) {
	httpRequestHistogram.With(prometheus.Labels{
		"method": method,
		"path":   path,
		"status": fmt.Sprintf("%d", status),
	}).Observe(time.Since(start).Seconds())
}

func init() {
	reg := prometheus.NewRegistry()

	reg.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		httpRequestHistogram,
	)
}

func RequestMetrics(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		t1 := time.Now()
		defer func() {
			httpRequestCountInc(r.Method, r.RequestURI, ww.Status(), t1)
		}()

		next.ServeHTTP(ww, r)
	}
	return http.HandlerFunc(fn)
}
