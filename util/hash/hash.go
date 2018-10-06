package hash

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
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

func Md5int32Array(v []uint32) []byte {
	arrayLen := len(v)
	bs := make([]byte, arrayLen*4)

	for i := 0; i < arrayLen; i++ {
		binary.LittleEndian.PutUint32(bs[i*4:], v[i])
	}
	hash := md5.Sum(bs)
	fmt.Printf("%x\n", bs)
	return hash[0:]
}
