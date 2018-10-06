package task_queue

type ITaskInterface interface {
	Add(task interface{}) error
	Get() (interface{}, error)
}
