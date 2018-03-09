package logger

import (
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/cihub/seelog"
)

func InitLogger(confPath string) error {
	defer log.Flush()
	hostname, _ := os.Hostname()
	template, err := ioutil.ReadFile(confPath)
	if err != nil {
		return err
	}
	data := string(template[:])
	config := fmt.Sprintf(data, hostname, hostname)
	logger, _ := log.LoggerFromConfigAsBytes([]byte(config))
	log.ReplaceLogger(logger)
	return nil
}
