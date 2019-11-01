package main

import (
	"log"

	"github.com/wangyihit/libmt/task_queue"
	"github.com/wangyihit/libmt/task_queue/redis_queue"
	"github.com/wangyihit/libmt/task_queue/task"
	"github.com/wangyihit/libmt/util/options"
	"github.com/wangyihit/libmt/util/options/redis_options"
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
