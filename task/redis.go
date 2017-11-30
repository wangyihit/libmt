package task

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

type RedisTaskManager struct {
	Host string
	Port int
	Addr string
	Key  string
}

func NewRedisTaskManager(host string, port int, key string) *RedisTaskManager {
	m := &RedisTaskManager{
		Host: host,
		Port: port,
		Addr: fmt.Sprintf("%s:%d", host, port),
		Key:  key,
	}
	return m
}

func (m *RedisTaskManager) GetTask() ([]byte, error) {
	c, err := redis.Dial("tcp", m.Addr)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	c.Send("RPOP", m.Key)
	c.Flush()
	data, err := redis.Bytes(c.Receive())
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (m *RedisTaskManager) AddTask(bytes []byte) error {
	c, err := redis.Dial("tcp", m.Addr)
	if err != nil {
		return nil
	}
	defer c.Close()
	c.Send("LPUSH", m.Key, bytes)
	c.Flush()
	_, err = redis.String(c.Receive())

	return nil
}
