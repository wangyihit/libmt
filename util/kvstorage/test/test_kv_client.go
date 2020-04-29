package main

import (
	"fmt"
	"time"

	"github.com/wangyihit/libmt/util/kvstorage"
)

func testKV() {
	var client kvstorage.IKvStoreClient
	var err error
	endpoint := "minio.mt"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	useSSl := false
	bucket := "images"
	imgName := "timg.jpeg"
	client = kvstorage.NewMinio(endpoint, accessKeyID, secretAccessKey, useSSl, bucket)
	err = client.Init()
	if err != nil {
		fmt.Printf("Client init failed, msg=%s\n", err.Error())
		return
	}
	exist, err := client.Exist(imgName)
	if err != nil {
		fmt.Printf("check image status failed, msg=%s\n", err.Error())
		return
	}
	fmt.Printf("image, Name=%s, exit=%+v\n", imgName, exist)
	content := "test data"
	testObjName := "test"
	err = client.Put(testObjName, []byte(content), kvstorage.ContentTypeText)
	if err != nil {
		fmt.Printf("put data error, msg=%s", err.Error())
	} else {
		fmt.Println("put object success")
	}
	bytes, err := client.Get(testObjName)
	if err != nil {
		fmt.Printf("get obj failed, msg=%s\n", err.Error())
	} else {
		fmt.Printf("get data: %s\n", string(bytes))
	}
	url, err := client.ShareUrl(testObjName, time.Second*87500)
	if err != nil {
		fmt.Printf("gen share url failed, msg=%s\n", err.Error())
	} else {
		fmt.Printf("gen url\n%s\n", url)
	}
}
func main() {
	testKV()
}
