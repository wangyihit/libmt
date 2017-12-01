package task

const (
	TaskIDTest     = 0
	TaskIDJson     = 1
	TaskIDJsonTest = 2
	TaskIDMax      = iota
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
	// TaskName() string
	// TaskTypeID() int
	FromBytes() error
	ToBytes() ([]byte, error)
}

type TaskManager interface {
	GetTask() ([]byte, error)
	AddTask([]byte) error
}
