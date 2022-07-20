package processor

import (
	"errors"

	"github.com/robertkrimen/otto"
	"github.com/wangyihit/thrift_idl/go/thrift_gen/mt/processor"
)

type Js struct {
	vm *otto.Otto
}

var _ Processor = (*Js)(nil)

func NewJs() *Js {
	vm := otto.New()
	return &Js{vm: vm}
}

func (c *Js) Name() string {
	return processor.JSTask
}

func (c *Js) Run(task *processor.TTask) (*processor.TTaskResult_, error) {
    taskData := task.GetTaskData()
	jsCmd := taskData.GetData()
	c.vm.Run(jsCmd)
	data, err := c.vm.Get("result")
	if err != nil {
		return nil, err
	}
	taskResult := NewTaskResult_()
	taskResult.Data, err = data.ToString()
	if err != nil {
		taskResult.TaskStatus = processor.TTaskStatus_FAILED
		return taskResult, err
	}
	if taskResult.GetData() == "Infinity" {
		taskResult.TaskStatus = processor.TTaskStatus_FAILED
		return taskResult, errors.New("js code error")
	}
	taskResult.TaskStatus = processor.TTaskStatus_SUCCESS
	return taskResult, nil
}
