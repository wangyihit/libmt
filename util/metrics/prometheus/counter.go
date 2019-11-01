package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

type LocalCounter struct {
	c prometheus.Counter
}

func (c *LocalCounter) Emit(n float64) {
	c.c.Add(n)
}
