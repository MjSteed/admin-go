package bootstrap

import (
	"os"

	"github.com/MjSteed/vue3-element-admin-go/common"
	"github.com/MjSteed/vue3-element-admin-go/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger() *zap.Logger {
	var level zapcore.Level
	var options []zap.Option
	switch common.Config.Logger.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
	if common.Config.Logger.ShowLine {
		options = append(options, zap.AddCaller())
	}
	// 初始化 zap
	return zap.New(getZapCore(level), options...)
}

// 扩展 Zap
func getZapCore(level zapcore.Level) zapcore.Core {
	var encoder zapcore.Encoder

	// 调整编码器默认配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// 设置编码器
	if common.Config.Logger.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	//同时输出至控制台和文件
	return zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(os.Stdout, getLogWriter()), level)
}

// 使用 lumberjack 作为日志写入器
func getLogWriter() zapcore.WriteSyncer {
	utils.CreateDir(common.Config.Logger.RootDir)
	file := &lumberjack.Logger{
		Filename:   common.Config.Logger.RootDir + "/" + common.Config.Logger.Filename,
		MaxSize:    common.Config.Logger.MaxSize,
		MaxBackups: common.Config.Logger.MaxBackups,
		MaxAge:     common.Config.Logger.MaxAge,
		Compress:   common.Config.Logger.Compress,
	}
	return zapcore.AddSync(file)
}
