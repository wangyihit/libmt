package task

const (
	TaskIDTest     = 0
	TaskIDJson     = 1
	TaskIDJsonTest = 2
	TaskIDMax      = iota
)

const (
	GetTaskSuccess = 0
	GetTaskFail    = iota
)

const (
	AddTaskSuccess = 0
	AddTaskFail    = iota
)

type TaskParser interface {
	FromBytes(bytes []byte, data interface{}) error
	ToBytes(data interface{}) ([]byte, error)
}

type TaskManager interface {
	GetTask() ([]byte, error)
	AddTask([]byte) error
}

type TaskHelper struct {
	taskManager TaskManager
	taskParser  TaskParser
}

func NewTaskHelper(taskManager TaskManager, taskParser TaskParser) *TaskHelper {
	helper := &TaskHelper{
		taskManager: taskManager,
		taskParser:  taskParser,
	}
	return helper
}

func (h *TaskHelper) GetTask(task interface{}) error {
	data, err := h.taskManager.GetTask()
	if err != nil {
		return err
	}
	err = h.taskParser.FromBytes(data, task)
	return err
}

func (h *TaskHelper) AddTask(task interface{}) error {
	bytes, err := h.taskParser.ToBytes(task)
	if err != nil {
		return err
	}
	err = h.taskManager.AddTask(bytes)
	return err
}
