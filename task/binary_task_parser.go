package task

import (
	"errors"

	"git.apache.org/thrift.git/lib/go/thrift"
)

// serilize task with thrift binary serilizer

type BinaryTaskParser struct {
	buffer       *thrift.TMemoryBuffer
	protocol     thrift.TProtocol
	serializer   *thrift.TSerializer
	deSerializer *thrift.TDeserializer
}

func NewBinaryTaskParser(buffer_size int) *BinaryTaskParser {
	t := thrift.NewTMemoryBufferLen(1024)
	p := thrift.NewTBinaryProtocolFactoryDefault().GetProtocol(t)

	tser := &thrift.TSerializer{
		Transport: t,
		Protocol:  p,
	}
	dser := &thrift.TDeserializer{
		Transport: t,
		Protocol:  p,
	}
	parser := &BinaryTaskParser{
		buffer:       t,
		protocol:     p,
		serializer:   tser,
		deSerializer: dser,
	}
	return parser
}

func (t *BinaryTaskParser) FromBytes(bytes []byte, task interface{}) error {
	t.buffer.Close() // resets underlying bytes.Buffer
	if tt, ok := task.(thrift.TStruct); ok {
		return t.deSerializer.Read(tt, bytes)
	}
	return errors.New("not binary task")
}

func (t *BinaryTaskParser) ToBytes(task interface{}) ([]byte, error) {
	t.buffer.Close() // resets underlying bytes.Buffer
	if tt, ok := task.(thrift.TStruct); ok {
		return t.serializer.Write(tt)
	}
	return nil, errors.New("not binary task")
}
