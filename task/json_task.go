package task

import (
	"encoding/json"
)

type JsonTask struct {
	Name string
	ID   int
}

func NewJsonTask() *JsonTask {
	t := &JsonTask{
		Name: "JsonTask",
		ID:   TaskIDJson,
	}
	return t
}

func (t *JsonTask) TaskName() string {
	return t.Name
}
func (t *JsonTask) TaskTypeID() int {
	return t.ID
}
func (t *JsonTask) FromBytes(bytes []byte) error {
	err := json.Unmarshal(bytes, t)
	if err != nil {
		return err
	}
	return nil
}
func (t *JsonTask) ToBytes() ([]byte, error) {
	return json.Marshal(t)
}
