package task_queue

import (
	"sync"
	"time"

	"go.uber.org/zap"
)

type Consumer struct {
	queue         ITaskQueue
	worker        IWorker
	keepRunning   bool
	sleepDuration int
	wg            *sync.WaitGroup
}

func NewConsumer(q ITaskQueue, worker IWorker, wg *sync.WaitGroup) *Consumer {
	c := &Consumer{
		queue:         q,
		worker:        worker,
		keepRunning:   true,
		sleepDuration: 1,
		wg:            wg,
	}
	return c
}
func (c *Consumer) Sleep() {
	time.Sleep(time.Second * time.Duration(c.sleepDuration))
	if c.sleepDuration < 64 {
		c.sleepDuration = c.sleepDuration * 2
	}
}
func (c *Consumer) Start() {
	if c.wg != nil {
		defer c.wg.Done()
	}
	for c.keepRunning == true {
		taskData, err := c.queue.Get()
		if err != nil {
			zap.L().Warn("Get task failed", zap.String("queue_info", c.queue.QueueInfo()))
			c.Sleep()
			continue
		}
		c.worker.Run(taskData)
		c.sleepDuration = 1
	}
}

func (c *Consumer) Stop() {
	c.keepRunning = false
}
