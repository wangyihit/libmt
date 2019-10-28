package main

import (
	"lib.mt/util/log"
	"lib.mt/util/log/loggers"
)

func main() {
	logger := loggers.NewDefault("DefaultMain")
	log.SetLogger(logger)
	log.Logger.Error("I am error")
	log.Logger.Info("I am info")
}
