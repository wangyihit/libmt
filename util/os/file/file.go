package file

import (
	"io/ioutil"
)

func Write(fileName string, bytes []byte) error {
	err := ioutil.WriteFile(fileName, bytes, 0644)
	return err
}

func Read(fileName string) ([]byte, error) {
	return ioutil.ReadFile(fileName)
}
