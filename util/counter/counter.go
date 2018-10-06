package counter

import "sync/atomic"

type Counter struct {
	num int64
}

func NewCounter() *Counter {
	c := &Counter{
		num: 0,
	}
	return c
}

func (c *Counter) Add(n int64) {
	atomic.AddInt64(&c.num, n)
}

func (c *Counter) Sub(n int64) {
	atomic.AddInt64(&c.num, n*-1)
}
