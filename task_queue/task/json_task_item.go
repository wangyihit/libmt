package task

import "encoding/json"

type Task struct {
	data []byte
}

func NewJsonTask(data []byte) *Task  {
	task := &Task{data:data}
	return task
}

func (t *Task)FromBytes(bytes []byte) error {
	t.data = bytes
	return json.Unmarshal(bytes, t)
}
func (t *Task) ToBytes() ([]byte, error) {
	return t.data, nil
}
