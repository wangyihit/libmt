package main

import (
	"fmt"

	"github.com/wangyihit/libmt/processor"
)

func main() {
	fmt.Println("run")
	jsProcessor := processor.NewJs()
	jsCmd := "var a = 5; var result = a * a;"
	res, err := jsProcessor.Run(jsCmd)
	resStr := res.(string)
	fmt.Printf("res=%s, err=%s", resStr, err)
}
