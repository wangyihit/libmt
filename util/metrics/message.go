package metrics

import "time"

type Timer struct {
	Name  string
	Start int64 // unix time stamp ms
	End   int64
}

func NewTimer(name string) *Timer {
	timeNow := time.Now().UnixNano() / 1000000
	t := &Timer{
		Name:  name,
		Start: timeNow,
		End:   timeNow,
	}
	return t
}

func (t *Timer) Stop() {
	t.End = time.Now().UnixNano() / 1000000
}

func (t *Timer) Ms() int64 {
	return t.End - t.Start
}

func (t *Timer) Reset() {
	t.Start = time.Now().UnixNano() / 1000000
	t.End = t.Start
}
