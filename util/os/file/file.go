package file

import (
	"io/ioutil"
	"os"
)

func Write(fileName string, bytes []byte) error {
	err := ioutil.WriteFile(fileName, bytes, 0644)
	return err
}

func Read(fileName string) ([]byte, error) {
	return ioutil.ReadFile(fileName)
}

func FileExist(fileName string) bool {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return false
	}
	return true
}