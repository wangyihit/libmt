package processor

import (
	"github.com/wangyihit/thrift_idl/go/thrift_gen/mt/processor"
)

type Processor interface {
	Run(data *processor.Task) (*processor.TaskResult_, error)
	Name() string
}

var NewTask = processor.NewTask

func NewTaskResult_() *processor.TaskResult_ {
	t := processor.NewTaskResult_()
	t.Extra = make(map[string]string)
	return t
}
