package bootstrap

import (
	"time"

	"github.com/MjSteed/vue3-element-admin-go/common"
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
		options = append(options, zap.AddStacktrace(level))
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
		options = append(options, zap.AddStacktrace(level))
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
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("[" + "2023-03-15 20:42:05.000" + "]"))
	}
	encoderConfig.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(l.String())
	}

	// 设置编码器
	if common.Config.Logger.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	return zapcore.NewCore(encoder, getLogWriter(), level)
}

// 使用 lumberjack 作为日志写入器
func getLogWriter() zapcore.WriteSyncer {
	file := &lumberjack.Logger{
		Filename:   common.Config.Logger.RootDir + "/" + common.Config.Logger.Filename,
		MaxSize:    common.Config.Logger.MaxSize,
		MaxBackups: common.Config.Logger.MaxBackups,
		MaxAge:     common.Config.Logger.MaxAge,
		Compress:   common.Config.Logger.Compress,
	}

	return zapcore.AddSync(file)
}
