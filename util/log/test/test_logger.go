package main

import (
	"libmt/util/log"
	"libmt/util/log/loggers"
)

func main() {
	logger := loggers.NewDefault("DefaultMain")
	log.SetLogger(logger)
	log.Logger.Error("I am error")
	log.Logger.Info("I am info")
}
