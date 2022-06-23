package main

import (
	"fmt"

	"github.com/wangyihit/libmt/processor"
)

func main() {
	fmt.Println("run")
	jsProcessor := processor.NewJs()
	jsCmd := "var a = 5; var result = a * a / 0;"
	task := processor.NewTask()
	task.Data = jsCmd
	res, err := jsProcessor.Run(task)
	resStr := res.GetData()
	fmt.Printf("res=%s, err=%s\n", resStr, err)
}
