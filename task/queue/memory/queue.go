package memory

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/wangyihit/libmt/task/queue"
)

type MemoryQueue struct {
	tasks chan interface{}
	size  int
}

var _ queue.ITaskQueue = (*MemoryQueue)(nil)

func NewMemoryQueue(size int) *MemoryQueue {
	q := &MemoryQueue{
		size:  size,
		tasks: make(chan interface{}, size),
	}
	return q
}

func (q *MemoryQueue) Add(task []byte) error {
	select {
	case q.tasks <- task:
		return nil
	default:
		return errors.New("Add task to queue failed")
	}
}

func (q *MemoryQueue) Get() ([]byte, error) {
	select {
	case task := <-q.tasks:
		return task.([]byte), nil
	default:
		return nil, errors.New("Get task failed")
	}
}

func (q *MemoryQueue) QueueInfo() string {
	return fmt.Sprintf("size=%d", q.size)
}
