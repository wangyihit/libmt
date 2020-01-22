package prometheus

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/wangyihit/libmt/util/metrics"
)

type LocalMetrics struct {
	Namespace string
	Subsystem string
}

func NewLocalMetrics(namespace string, subsystem string) *LocalMetrics {
	return &LocalMetrics{
		Namespace: namespace,
		Subsystem: subsystem,
	}
}

func (m *LocalMetrics) NewCounter(name string, labels map[string]string, help string) metrics.ICounter {
	opts := prometheus.CounterOpts{
		Namespace:   m.Namespace,
		Subsystem:   m.Subsystem,
		Name:        name,
		Help:        help,
		ConstLabels: labels,
	}
	counter := promauto.NewCounter(opts)
	localCounter := &LocalCounter{c: counter}
	// prometheus.MustRegister(counter)
	return localCounter
}

var Metrics *LocalMetrics
var defaultTimerObjectives = map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001}

func (m *LocalMetrics) NewSummary(name string, labels []string, Objectives map[float64]float64, help string) metrics.ISummary {
	if Objectives == nil {
		Objectives = defaultTimerObjectives
	}
	v := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       name,
			Help:       help,
			Objectives: Objectives,
		},
		labels,
	)
	prometheus.MustRegister(v)
	localTimer := NewLocalTimer(v)
	return localTimer
}

func Start(addr string) {
	http.Handle("/metrics", promhttp.Handler())
	_ = http.ListenAndServe(addr, nil)
}

func InitGlobal(namespace string, subsystem string) {
	Metrics = NewLocalMetrics(namespace, subsystem)
}
