package serializer

import (
	"context"

	"github.com/apache/thrift/lib/go/thrift"
)

func BytesToThriftObject(bytes []byte, i thrift.TStruct) error {
	deSerializer := thrift.NewTDeserializer()
	return deSerializer.Read(i, bytes)
}

func ThriftObjectToBytes(i thrift.TStruct) ([]byte, error) {

	memoryBuffer := thrift.NewTMemoryBufferLen(1024 * 1024 * 10)
	protocol := thrift.NewTBinaryProtocolFactoryDefault().GetProtocol(memoryBuffer)
	serializer := &thrift.TSerializer{
		Transport: memoryBuffer,
		Protocol:  protocol,
	}
	serializer.Transport.Close()
	return serializer.Write(context.TODO(), i)

}
