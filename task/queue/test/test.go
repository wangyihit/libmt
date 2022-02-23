package main

import (
	"log"

	"github.com/wangyihit/libmt/task/queue"
	redisQueue "github.com/wangyihit/libmt/task/queue/redis"
	"github.com/wangyihit/libmt/task/queue/task"
	"github.com/wangyihit/libmt/util/options"
	redisOpeions "github.com/wangyihit/libmt/util/options/redis"
	"gopkg.in/redis.v3"
)

func main() {
	redisHost := "172.17.10.140"
	redisPort := int64(6379)
	redisPassword := "password"
	redisOptions := redisOpeions.NewOptions(redisHost, redisPort, redisPassword, 0)
	options.Init()
	log.Printf("Redis option: %+v\n", redisOptions)
	var testQueue queue.ITaskQueue
	taskData := `{"name":"wangyi","info":"empty"}`
	taskItem := task.NewJsonTask([]byte(taskData))
	// host string, port int64, password string, tag string, queueName string
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	testQueue = redisQueue.NewRedisQueue(redisClient, "wangyi_test", "test_queue")
	bytes, _ := taskItem.ToBytes()
	err := testQueue.Add(bytes)
	log.Printf("msg=add_task, msg=%+v\n", err)
}
