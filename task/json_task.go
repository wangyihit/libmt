package task

import (
	"encoding/json"
)

type JsonTaskParser struct {
}

func NewJsonTaskParser() *JsonTaskParser {
	t := &JsonTaskParser{}
	return t
}

func (t *JsonTaskParser) FromBytes(bytes []byte, task interface{}) error {
	err := json.Unmarshal(bytes, task)
	if err != nil {
		return err
	}
	return nil
}
func (t *JsonTaskParser) ToBytes(task interface{}) ([]byte, error) {
	return json.Marshal(task)
}
