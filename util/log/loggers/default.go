package loggers

import (
	"fmt"
	"log"
	"os"
)
var logDepth = 2

type Default struct {
	info *log.Logger
	warn *log.Logger
	error *log.Logger
}

func NewDefault(loggerName string) *Default {
	info := log.New(os.Stdout, loggerName, log.Llongfile)
	err := log.New(os.Stdout, loggerName, log.Llongfile)
	l := &Default{
		info:  info,
		warn:  err,
		error: err,
	}
	return l
}

func (l *Default) Info(args ...interface{}) {
	_ = l.info.Output(logDepth, fmt.Sprintf(args[0].(string) + "\n", args[1:]...))
}

func (l *Default) Warn(args ...interface{}) {
	_ = l.warn.Output(logDepth, fmt.Sprintf(args[0].(string) + "\n", args[1:]...))
}

func (l *Default) Error(args ...interface{}) {
	_ = l.error.Output(logDepth, fmt.Sprintf(args[0].(string) + "\n", args[1:]...))
}

