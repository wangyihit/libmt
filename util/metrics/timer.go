package metrics

import (
	"time"
)

type Timer struct {
	start time.Time
	m     ISummary
}

func NewTimer(metrics ISummary) *Timer {
	t := time.Now()
	timer := &Timer{
		start: t,
		m:     metrics,
	}
	return timer
}

func (t *Timer) duration() time.Duration {
	return time.Since(t.start)
}

func (t *Timer) Emit(c float64, labels []string) {
	t.m.Emit(c, labels)
}

func (t *Timer) EmitLater(labels []string) {
	ns := t.duration().Nanoseconds()
	t.m.Emit(float64(ns), labels)
}
