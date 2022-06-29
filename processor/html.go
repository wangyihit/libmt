package processor

import (
	"github.com/wangyihit/thrift_idl/go/thrift_gen/mt/processor"
)

// var NewHtmlProcessor = processor.NewHtmlBlock
// var NewHtmlEXtractResult_ = processor.NewHtmlEXtractResult_
// var NewItemField = processor.NewItemField

type HtmlExtractTask processor.HtmlExtractTask

type HtmlProcessor struct {
}

func NewHtmlProcesor() *HtmlProcessor {
	p := &HtmlProcessor{}
	return p
}



func (c *HtmlProcessor) Name() string {
	return "HtmlProcessor"
}

func ParseTask(data string) * processor.HtmlExtractTask {
    task := processor.NewHtmlExtractTask()
    return task
}

func (c *HtmlProcessor) Run(task *processor.Task) (*processor.TaskResult_, error) {
	taskData := task.GetData()
	taskResult := NewTaskResult_()
    taskResult.Data = taskData
	return taskResult, nil
}

