package trace

import (
	"github.com/cihub/seelog"
)

func TracePanic(i interface{}) {
	if err := recover(); err != nil {
		seelog.Warnf("Panic, error=%+v, info=%+v\n", err, i)
	}
}
