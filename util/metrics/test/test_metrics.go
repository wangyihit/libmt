package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/wangyihit/libmt/util/metrics"
	"github.com/wangyihit/libmt/util/metrics/prometheus"
)

func testCounter(m metrics.IMetrics) {
	c := m.NewCounter("test_counter_0001", nil, "test counter 001")
	for {
		c.Emit(1)
		time.Sleep(1 * time.Second)
	}
}
func testSummary(m metrics.IMetrics) {
	ls := []string{"l1", "l2"}
	s := m.NewSummary("test_summary_001", ls, nil, "test summary 001")
	t := metrics.NewTimer(s)
	for {
		r := rand.ExpFloat64()
		t.Emit(r, ls)
		time.Sleep(1)
	}
}
func main() {
	m := prometheus.NewLocalMetrics("crawler", "reg_server")

	go testCounter(m)
	go testSummary(m)
	// Expose the registered metrics via HTTP.
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":2112", nil))
}
