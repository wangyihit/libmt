package main

import (
	"fmt"

	"github.com/wangyihit/libmt/util/thrift/serializer"
	"github.com/wangyihit/libmt/util/thrift/serializer/test/gen-go/test"
)

func _main() {
	t := test.NewTest()
	t.Name = "NewName"
	bytes, err := serializer.ThriftObjectToBytes(t)
	if err != nil {
		fmt.Print(err)
		return
	}
	t2 := test.NewTest()
	err = serializer.BytesToThriftObject(bytes, t2)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("%+v\n", t2)
	}
}

func main() {
	_main()
}
