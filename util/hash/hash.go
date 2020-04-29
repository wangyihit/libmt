package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

var _ = fmt.Println

func Md5Int64(v uint64) []byte {
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, v)
	hash := md5.Sum(bs)
	return hash[0:]
}

func Md5Int32(v uint32) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, v)
	hash := md5.Sum(bs)
	return hash[0:]
}

func Md5Int32Array(v []uint32) []byte {
	arrayLen := len(v)
	bs := make([]byte, arrayLen*4)

	for i := 0; i < arrayLen; i++ {
		binary.LittleEndian.PutUint32(bs[i*4:], v[i])
	}
	hash := md5.Sum(bs)
	fmt.Printf("%x\n", bs)
	return hash[0:]
}

func B64ToBytes(s []byte) ([]byte, error) {
	// func (enc *Encoding) Decode(dst, src []byte) (n int, err error)
	var b []byte
	_, err := base64.StdEncoding.Decode(b, s)
	return b, err
}

func BytesToB64(s []byte) []byte {
	var b []byte
	base64.StdEncoding.Encode(b, s)
	return b
}

func Sha1Hex(b []byte) string {
	h := sha1.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}
func Sha1HexFile(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", nil
	}
	defer f.Close()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", nil
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
func MD5Hex(b []byte) string {
	h := md5.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

func MD5HexFile(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", nil
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", nil
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
