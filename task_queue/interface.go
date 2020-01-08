package task_queue

type ITaskQueue interface {
	Add(task []byte) error
	Get() ([]byte, error)
	QueueInfo() string
}

type ITaskItem interface {
	FromBytes(bytes []byte) error
	ToBytes() ([]byte, error)
}

type IWorker interface {
	Run(task []byte)
}
