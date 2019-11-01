package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
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
