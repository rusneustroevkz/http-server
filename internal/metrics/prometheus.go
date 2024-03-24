package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	namespace = "test"
	subsystem = "backend_example"
)

var (
	httpRequestCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "http_request_total",
		Help:      "Http request total count",
	}, []string{"method", "code", "path"})
)

func httpRequestCountInc(method, path string, code int64) {
	httpRequestCount.With(prometheus.Labels{
		"method": method,
		"path":   path,
		"code":   fmt.Sprintf("%d", code),
	}).Inc()
}

func init() {
	reg := prometheus.NewRegistry()

	reg.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		httpRequestCount,
	)
}
