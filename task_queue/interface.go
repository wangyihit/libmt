package task_queue

type ITaskInterface interface {
	Add(task []byte) error
	Get() ([]byte, error)
}

type ITaskItem interface {
	FromBytes(bytes []byte) error
	ToBytes() ([]byte, error)
}
