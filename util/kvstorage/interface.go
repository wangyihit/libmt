package kvstorage

import "time"

type IKvStoreClient interface {
	Init() error
	Put(objName string, data []byte, contentType string) error
	Get(objName string) ([]byte, error)
	Exist(objName string) (bool, error)
	ShareUrl(objName string, expires time.Duration) (string, error)
}
