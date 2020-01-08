package main

import (
	"github.com/wangyihit/libmt/util/logger"
	"go.uber.org/zap"
)

func main() {
	logger.InitLogger("./logs/test_log", logger.LogLevelInfo)
	zap.L().Info("test info", zap.String("key", "v"))
	zap.L().Error("test err", zap.String("k", "v"))
}
