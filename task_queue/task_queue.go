package task_queue

import (
	"github.com/pkg/errors"
)

type MemoryQueue struct {
	tasks chan interface{}
	size  int
}

func NewMemoryQueue(size int) *MemoryQueue {
	q := &MemoryQueue{
		size:  size,
		tasks: make(chan interface{}, size),
	}
	return q
}

func (q *MemoryQueue) Add(task interface{}) error {
	select {
	case q.tasks <- task:
		return nil
	default:
		return errors.New("Add task to queue failed")
	}
}

func (q *MemoryQueue) Get() (interface{}, error) {
	select {
	case task := <-q.tasks:
		return task, nil
	default:
		return nil, errors.New("Get task failed")
	}
}

func CreateMemoryTaskQueues(queueCount int, queueSize int) []ITaskInterface {
	qs := make([]ITaskInterface, queueCount)
	for i := 0; i < queueCount; i++ {
		qs[i] = NewMemoryQueue(queueSize)
	}
	return qs
}
