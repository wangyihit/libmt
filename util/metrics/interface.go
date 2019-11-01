package metrics

type IMetrics interface {
	NewCounter(name string, labels map[string]string, help string) ICounter
	// for timer
	NewSummary(name string, labels []string, Objectives map[float64]float64, shelp string) ISummary
}

type ICounter interface {
	Emit(c float64)
}

type ISummary interface {
	Emit(c float64, labels []string)
}
