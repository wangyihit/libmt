package exepath

import (
	"os"
	"path/filepath"
)

func ExeDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	return dir
}
