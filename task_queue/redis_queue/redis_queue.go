package redis_queue

import (
	"fmt"

	"gopkg.in/redis.v3"
)

const queueNamePrefix = "task_queue"

type RedisQueue struct {
	port int64
	host string
	password string
	db int64
	tag string
	queueName string
	realQueueName string
	redisClient *redis.Client
}

func realQueueName(tag string , name string) string {
	n := fmt.Sprintf("%s::%s::%s", queueNamePrefix, tag, name)
	return n
}

func NewRedisQueue(host string, port int64, password string, tag string, queueName string) *RedisQueue{
	// tag :: service tag, queue name prefix

	q := &RedisQueue{
		port:      port,
		host:      host,
		password: password,
		queueName: queueName,
		realQueueName:realQueueName(tag, queueName),
		db: 0,
		redisClient:nil,
	}
	addr := fmt.Sprintf("%s:%d", host, port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       q.db,
	})
	q.redisClient = client
	return q
}

func (q *RedisQueue) Add(task []byte) error {
	cmd := q.redisClient.LPush(q.realQueueName, string(task))
	return cmd.Err()
}

func (q *RedisQueue) Get() ([]byte, error) {
	cmd := q.redisClient.RPop(q.realQueueName)
	return cmd.Bytes()
}
