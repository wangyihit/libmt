package compress

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io"
)

func UnGzipData(data []byte) (resData []byte, err error) {
	b := bytes.NewBuffer(data)
	var r io.Reader
	r, err = gzip.NewReader(b)
	if err != nil {
		return nil, errors.New("read data failed")
	}

	var resB bytes.Buffer
	_, err = resB.ReadFrom(r)
	if err != nil {
		return nil, errors.New("copy data failed")
	}

	resData = resB.Bytes()
	return resData, nil
}

func GZipData(data []byte) (compressedData []byte, err error) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)

	_, err = gz.Write(data)
	if err != nil {
		return nil, errors.New("write failed")
	}

	if err = gz.Flush(); err != nil {
		return nil, errors.New("flush failed")
	}
	if err = gz.Close(); err != nil {
		return nil, errors.New("close failed")
	}
	compressedData = b.Bytes()
	return compressedData, nil
}
