package processor

import (
	"github.com/wangyihit/thrift_idl/go/thrift_gen/mt/processor"
)

type Processor interface {
	Run(data *processor.Task) (*processor.TaskResult_, error)
	Name() string
}

var NewTask = processor.NewTask
var NewHtmlProcessor = processor.NewHtmlBlock
var NewHtmlEXtractResult_ = processor.NewHtmlEXtractResult_
var NewItemField = processor.NewItemField

func NewTaskResult_() *processor.TaskResult_ {
	t := processor.NewTaskResult_()
	t.Extra = make(map[string]string)
	return t
}
