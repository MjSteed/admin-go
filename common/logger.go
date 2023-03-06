package common

import "go.uber.org/zap"

var LOG *zap.Logger

func init() {
	logger, _ := zap.NewDevelopment()
	LOG = logger
	defer LOG.Sync()
	LOG.Info("日志组件初始化成功")
	LOG.Sugar().Info("日志组件初始化成功", "测试", "test")
}
