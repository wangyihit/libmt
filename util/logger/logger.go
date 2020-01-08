package logger

import (
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/cihub/seelog"
)

func InitLogger(confPath string) error {
	defer log.Flush()
	hostName, _ := os.Hostname()
	confTemplate, err := ioutil.ReadFile(confPath)
	if err != nil {
		return err
	}
	data := string(confTemplate[:])
	config := fmt.Sprintf(data, hostName, hostName)
	logger, _ := log.LoggerFromConfigAsBytes([]byte(config))
	log.ReplaceLogger(logger)
	return nil
}
