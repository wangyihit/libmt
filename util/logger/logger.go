package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	LogLevelDebug = iota - 1
	LogLevelInfo
	LogLevelError
)

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
func InitLogger(dir string, fileName string, loglevel int) error {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}
	logPath := fmt.Sprintf("%s/%s", dir, fileName)
	var level zapcore.Level
	switch loglevel {
	case LogLevelDebug:
		level = zap.DebugLevel
	case LogLevelInfo:
		level = zap.InfoLevel
	case LogLevelError:
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	ll := initLogger(logPath, level)

	if level < zap.ErrorLevel {
		errorLogger := initLogger(logPath+".err", zap.ErrorLevel)
		core := zapcore.NewTee(
			ll.Core(), errorLogger.Core(),
		)
		zap.ReplaceGlobals(zap.New(core))
		return nil
	}
	zap.ReplaceGlobals(ll)
	return nil
}
func initLogger(logPath string, level zapcore.Level) *zap.Logger {
	hook := lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    128, // megabytes
		MaxBackups: 300,
		MaxAge:     7, // days
		LocalTime:  true,
		Compress:   false,
	}

	cfg := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(level),
		OutputPaths: []string{logPath},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "M",

			LevelKey:    "L",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "T",
			EncodeTime: TimeEncoder,

			CallerKey:    "C",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	ll, _ := cfg.Build()
	w := zapcore.AddSync(&hook)
	ll.WithOptions(
		zap.WrapCore(
			func(zapcore.Core) zapcore.Core {
				return zapcore.NewCore(
					zapcore.NewConsoleEncoder(cfg.EncoderConfig),
					w,
					level)
			}))
	return ll
}
