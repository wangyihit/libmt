package task

const (
	TaskIDTest = 0
	TaskIDMax  = iota
)

const (
	GetTaskSuccess = 0
	GetTaskFail    = iota
)

const (
	AddTaskSuccess = 0
	AddTaskFail    = iota
)

type Task interface {
	TaskName() string
	TaskTypeID() int
	FromBytes() error
	ToBytes() ([]byte, error)
}

type TaskManager interface {
	GetTask() ([]byte, error)
	AddTask([]byte) error
}

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
