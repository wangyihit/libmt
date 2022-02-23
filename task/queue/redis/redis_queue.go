package redis

import (
	"fmt"
	"sync"

	"github.com/wangyihit/libmt/task/queue"
	"gopkg.in/redis.v3"
)

const queueNamePrefix = "task_queue"

type RedisQueue struct {
	db          int64
	tag         string
	queueName   string
	redisKey    string
	redisClient *redis.Client
	muRedis     sync.Mutex
}

var _ queue.ITaskQueue = (*RedisQueue)(nil)

func redisKey(tag string, name string) string {
	n := fmt.Sprintf("%s::%s::%s", queueNamePrefix, tag, name)
	return n
}

func NewRedisQueue(redisClient *redis.Client, tag string, queueName string) *RedisQueue {
	// tag :: service tag, queue name prefix

	q := &RedisQueue{
		queueName:   queueName,
		tag:         tag,
		redisKey:    redisKey(tag, queueName),
		db:          0,
		redisClient: redisClient,
	}
	return q
}

func (q *RedisQueue) Add(task []byte) error {
	q.muRedis.Lock()
	cmd := q.redisClient.LPush(q.redisKey, string(task))
	q.muRedis.Unlock()
	return cmd.Err()
}

func (q *RedisQueue) Get() ([]byte, error) {
	q.muRedis.Lock()
	cmd := q.redisClient.RPop(q.redisKey)
	q.muRedis.Unlock()
	return cmd.Bytes()
}

func (q *RedisQueue) QueueInfo() string {
	info := fmt.Sprintf("%+v", q)
	return info
}
