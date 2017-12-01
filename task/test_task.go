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

func NewJsonTestTask(data string) *JsonTask {
	test_task := &JsonTestTask{
		Name:  "Json Test Task",
		ID:    TaskIDJsonTest,
		SData: data,
	}
	json_task := &JsonTask{
		Data: test_task,
	}
	return json_task
}

func RunJsonTestTask() {

	testTask := NewJsonTestTask("I am data.")
	redisTaskManager := NewRedisTaskManager("127.0.0.1", 6379, "test_task_queue")
	b, _ := testTask.ToBytes()
	fmt.Printf("task bytes: %s", b)
	redisTaskManager.AddTask(b)
	data, _ := redisTaskManager.GetTask()
	fmt.Printf("get task, data=%s", data)
}
