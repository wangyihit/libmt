package kvstorage

import (
	"bytes"
	"io"
	"net/url"
	"time"

	"github.com/wangyihit/libmt/util/hash"

	"github.com/pkg/errors"

	"github.com/minio/minio-go/v6"
)

const (
	ContentTypeText = "text/plain"
	ContentTypeJPG  = "image/jpeg"
	ContentTypePNG  = "image/png"
)

type Minio struct {
	endpoint        string
	accessKeyID     string
	secretAccessKey string
	useSSl          bool
	bucket          string
	client          *minio.Client
}

func NewMinio(endpoint string, accessKeyID string, secretAccessKey string, useSSl bool, bucket string) *Minio {
	m := &Minio{
		endpoint:        endpoint,
		accessKeyID:     accessKeyID,
		secretAccessKey: secretAccessKey,
		useSSl:          useSSl,
		bucket:          bucket,
		client:          nil,
	}
	return m
}

func (m *Minio) contentType(s string) string {
	if s == "" {
		return "application/octet-stream"
	}
	return s
}

func (m *Minio) Init() error {
	client, err := minio.New(m.endpoint, m.accessKeyID, m.secretAccessKey, m.useSSl)
	if err == nil {
		m.client = client
	}
	return err
}

func (m *Minio) Put(objName string, data []byte, contentType string, checkObjName bool) error {
	if checkObjName == true {
		exist, err := m.Exist(objName)
		if err != nil {
			return err
		}
		if exist == true {
			return nil
		}
	}
	reader := bytes.NewReader(data)
	_, err := m.client.PutObject(m.bucket, objName, reader, int64(len(data)), minio.PutObjectOptions{ContentType: contentType})
	return err
}

func (m *Minio) PutWithSha1Name(data []byte, contentType string, checkObjName bool) error {
	objName := hash.Sha1Hex(data)
	return m.Put(objName, data, contentType, checkObjName)
}

func (m *Minio) PutFile(objectName string, filePath string, contentType string, checkObjName bool) error {
	if checkObjName == true {
		exist, err := m.Exist(objectName)
		if err != nil {
			return err
		}
		if exist == true {
			return nil
		}
	}
	_, err := m.client.FPutObject(m.bucket, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	return err
}

func (m *Minio) PutFileWithSha1Name(filePath string, contentType string, checkObjName bool) error {
	md5Hex, err := hash.MD5HexFile(filePath)
	if err != nil {
		return err
	}
	return m.PutFile(md5Hex, filePath, contentType, checkObjName)
}

func (m *Minio) Exist(objName string) (bool, error) {
	_, err := m.client.StatObject(m.bucket, objName, minio.StatObjectOptions{})
	if err != nil {
		errResponse := minio.ToErrorResponse(err)
		if errResponse.Code == "NoSuchKey" {
			return false, nil
		} else {
			return false, errors.New(errResponse.Code)
		}
	}
	return true, nil
}

func (m *Minio) Get(objName string) ([]byte, error) {
	obj, err := m.client.GetObject(m.bucket, objName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	objStat, err := obj.Stat()
	if err != nil {
		return nil, err
	}
	var bytes []byte
	bytes = make([]byte, objStat.Size)
	_, err = obj.Read(bytes)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return bytes, nil
}

func (m *Minio) GetFile(objName string, filePath string) error {
	return m.client.FGetObject(m.bucket, objName, filePath, minio.GetObjectOptions{})
}

func (m *Minio) ShareUrl(objName string, expires time.Duration) (string, error) {
	reqParams := make(url.Values)
	shareUrl, err := m.client.PresignedGetObject(m.bucket, objName, expires, reqParams)
	if err != nil {
		return "", err
	}
	return shareUrl.String(), nil
}
