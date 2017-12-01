package task

import (
	"encoding/json"
)

type JsonTask struct {
	Data interface{}
}

func NewJsonTask(data interface{}) *JsonTask {
	t := &JsonTask{
		Data: data,
	}
	return t
}

func (t *JsonTask) FromBytes(bytes []byte) error {
	err := json.Unmarshal(bytes, t.Data)
	if err != nil {
		return err
	}
	return nil
}
func (t *JsonTask) ToBytes() ([]byte, error) {
	return json.Marshal(t.Data)
}
