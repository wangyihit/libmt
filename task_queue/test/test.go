package main

import (
	"log"

	"lib.mt/task_queue"
	"lib.mt/task_queue/redis_queue"
	"lib.mt/task_queue/task"
	"lib.mt/util/options"
	"lib.mt/util/options/redis_options"
)

func main() {
	redisHost := "172.17.10.140"
	redisPort := int64(6379)
	redisPassword := "password"
	redisOptions := redis_options.NewOptions(redisHost, redisPort, redisPassword, true)
	options.Init()
	log.Printf("Redis option: %+v\n", redisOptions)
	var queue task_queue.ITaskInterface
	taskData := `{"name":"wangyi","info":"empty"}`
	taskItem := task.NewJsonTask([]byte(taskData))
	// host string, port int64, password string, tag string, queueName string
	queue = redis_queue.NewRedisQueue(redisOptions.Host, redisOptions.Port, redisOptions.Password,
		"wangyi_test", "test_queue")
	bytes, _ := taskItem.ToBytes()
	err := queue.Add(bytes)
	log.Printf("msg=add_task, msg=%+v\n", err)
}
