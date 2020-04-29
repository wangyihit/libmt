package kvstorage

import "time"

type IKvStoreClient interface {
	Init() error
	Put(objName string, data []byte, contentType string, checkObjName bool) error
	PutWithSha1Name(data []byte, contentType string, checkObjName bool) error
	PutFile(objectName string, filePath string, contentType string, checkObjName bool) error
	PutFileWithSha1Name(filePath string, contentType string, checkObjName bool) error
	Get(objName string) ([]byte, error)
	GetFile(objName string, filePath string) error
	Exist(objName string) (bool, error)
	ShareUrl(objName string, expires time.Duration) (string, error)
}
