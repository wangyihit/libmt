package processor

import (
	"github.com/robertkrimen/otto"
)

type Js struct {
	vm *otto.Otto
}

var _ Processor = (*Js)(nil)

func NewJs() *Js {
	vm := otto.New()
	return &Js{vm: vm}
}

func (c *Js) Name() string {
	return "Js"
}

func (c *Js) Run(data Variant) (Variant, error) {
	jsCmd := data.(string)
	c.vm.Run(jsCmd)
	res, err := c.vm.Get("result")
	if err != nil {
		return "", err
	}
	return res.ToString()
}
