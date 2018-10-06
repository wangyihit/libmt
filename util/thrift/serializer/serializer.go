package serializer

import "git.apache.org/thrift.git/lib/go/thrift"

func BytesToThriftObject(bytes []byte, i thrift.TStruct) error {

	memoryBuffer := thrift.NewTMemoryBufferLen(1024 * 1024 * 10)
	protocol := thrift.NewTBinaryProtocolFactoryDefault().GetProtocol(memoryBuffer)
	deSerializer := &thrift.TDeserializer{
		Transport: memoryBuffer,
		Protocol:  protocol,
	}
	deSerializer.Transport.Close()
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
	return serializer.Write(i)
}
