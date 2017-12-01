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

func (t *JsonTaskParser) FromBytes(bytes []byte, data interface{}) error {
	err := json.Unmarshal(bytes, data)
	if err != nil {
		return err
	}
	return nil
}
func (t *JsonTaskParser) ToBytes(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}
