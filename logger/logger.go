package logger

import (
	"bluebell_rebuild/setting"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Init(cfg setting.LogConfig, mode string) error {
	// 1. 创建日志文件写入器
	writeSyncer := getLogWriter(cfg)

	// 2. 创建JSON日志格式
	encoder := getEncoder()

	// 3. 解析配置中的日志等级
	level := new(zapcore.Level)

	if err := level.UnmarshalText([]byte(cfg.Level)); err != nil {
		return err
	}

	// 4. 创建写入日志文件的核心
	fileCore := zapcore.NewCore(
		encoder,
		writeSyncer,
		level,
	)

	var core zapcore.Core

	if mode == "dev" {
		// 开发模式下，控制台使用更容易阅读的格式
		consoleEncoder := zapcore.NewConsoleEncoder(
			zap.NewDevelopmentEncoderConfig(),
		)

		consoleCore := zapcore.NewCore(
			consoleEncoder,
			zapcore.Lock(os.Stdout),
			zapcore.DebugLevel,
		)

		// 同时输出到日志文件和控制台
		core = zapcore.NewTee(fileCore, consoleCore)
	} else {
		// 生产模式只写入日志文件
		core = fileCore
	}

	// 5. 根据core创建Logger，并记录代码位置
	lg := zap.New(core, zap.AddCaller())

	// 6. 设置成全局Logger
	zap.ReplaceGlobals(lg)

	zap.L().Info("logger init success")

	return nil
}

func getLogWriter(cfg setting.LogConfig) zapcore.WriteSyncer {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.MaxSize,
		MaxAge:     cfg.MaxAge,
		MaxBackups: cfg.MaxBackups,
	}
	return zapcore.AddSync(lumberjackLogger) //返回一个文件写入器
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	// 使用容易阅读的时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 时间字段名
	encoderConfig.TimeKey = "time"
	// 日志等级使用大写
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 代码位置只显示短路径
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig) //返回日志格式
}
