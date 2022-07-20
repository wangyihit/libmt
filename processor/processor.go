package processor

import (
	"github.com/wangyihit/thrift_idl/go/thrift_gen/mt/processor"
)

type Processor interface {
	Run(data *processor.TTask) (*processor.TTaskResult_, error)
	Name() string
}

var NewTask = processor.NewTTask


func NewTaskResult_() *processor.TTaskResult_ {
	t := processor.NewTTaskResult_()
	t.Extra = make(map[string]string)
    t.TaskStatus = processor.TTaskStatus_SUCCESS
	return t
}

