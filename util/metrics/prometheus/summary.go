package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/wangyihit/libmt/util/metrics"
)

type LocalSummary struct {
	s *prometheus.SummaryVec
}

func NewLocalTimer(s *prometheus.SummaryVec) *LocalSummary {
	return &LocalSummary{s: s}
}

func (c *LocalSummary) NewInstance(labels []string) *metrics.Timer {
	m := metrics.NewTimer(c)
	return m
}

func (c *LocalSummary) Emit(n float64, labels []string) {
	if labels == nil {
		labels = []string{}
	}
	c.s.WithLabelValues(labels...).Observe(n)
}
