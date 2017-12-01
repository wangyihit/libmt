package task

import (
	"fmt"
)

type TestTask struct {
	Data string
}

func NewTestTask(data string) *TestTask {
	t := &TestTask{
		Data: data,
	}
	return t
}

func (_ *TestTask) TaskName() string {
	return "TestTask"
}
func (_ *TestTask) TaskTypeID() int {
	return TaskIDTest
}
func (t *TestTask) FromBytes(bytes []byte) error {
	t.Data = string(bytes)
	return nil
}
func (t *TestTask) ToBytes() ([]byte, error) {
	return []byte(t.Data), nil
}

type JsonTestTask struct {
	Name  string
	ID    int
	SData string
}

func NewJsonTestTask(data string) *JsonTestTask {
	test_task := &JsonTestTask{
		Name:  "Json Test Task",
		ID:    TaskIDJsonTest,
		SData: data,
	}

	return test_task
}

func RunJsonTestTask() {

	testTask := NewJsonTestTask("I am data.")
	taskManager := NewRedisTaskManager("127.0.0.1", 6379, "test_task_queue")
	taskPaser := NewJsonTaskParser()
	task_helper := NewTaskHelper(taskManager, taskPaser)
	task_helper.AddTask(testTask)
	t := new(JsonTestTask)
	task_helper.GetTask(t)
	fmt.Printf("get task, data=%+v", t)
}
