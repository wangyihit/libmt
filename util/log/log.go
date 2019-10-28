package log

var Logger ILogger

func SetLogger(logger ILogger) {
	Logger = logger
}
