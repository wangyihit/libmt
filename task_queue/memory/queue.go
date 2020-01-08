package memory

import (
	"fmt"

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

func (q *MemoryQueue) QueueInfo() string {
	return fmt.Sprintf("size=%d", q.size)
}
