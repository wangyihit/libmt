package log

type ILogger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}
