package logger

import (
	"bluebell_rebuild/setting"

	"go.uber.org/zap"
)

func Init(cfg setting.LogConfig, mode string) error {
	var (
		lg  *zap.Logger
		err error
	)

	// 开发模式：日志更适合人阅读
	if mode == "dev" {
		lg, err = zap.NewDevelopment()
	} else {
		// 生产模式：日志以 JSON 形式输出
		lg, err = zap.NewProduction()
	}

	if err != nil {
		return err
	}

	// 将lg 替换为Zap的全局logger
	zap.ReplaceGlobals(lg)

	zap.L().Info(
		"logger init success",
		zap.String("level", cfg.Level),
		zap.String("filename", cfg.Filename),
		zap.Int("max_size", cfg.MaxSize),
		zap.Int("max_age", cfg.MaxAge),
		zap.Int("max_backups", cfg.MaxBackups),
	)

	return nil
}
