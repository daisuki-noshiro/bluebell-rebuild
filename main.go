package main

import (
	"fmt"
	"log"

	"bluebell_rebuild/logger"
	"bluebell_rebuild/router"
	"bluebell_rebuild/setting"

	"go.uber.org/zap"
)

func main() {
	// 1. 读取配置文件
	if err := setting.Init("./conf/config.yaml"); err != nil {
		log.Fatal("读取配置文件失败：", err)
	}

	// 2. 初始化日志
	if err := logger.Init(setting.Conf.Log, setting.Conf.Mode); err != nil {
		log.Fatal("初始化日志失败：", err)
	}

	// 程序退出前刷新尚未写出的日志
	defer zap.L().Sync()

	zap.L().Info(
		"bluebell server starting",
		zap.String("name", setting.Conf.Name),
		zap.String("mode", setting.Conf.Mode),
		zap.Int("port", setting.Conf.Port),
	)

	// 3. 注册路由
	r := router.SetupRouter()

	// 4. 启动服务器
	addr := fmt.Sprintf(":%d", setting.Conf.Port)

	if err := r.Run(addr); err != nil {
		zap.L().Fatal(
			"server start failed",
			zap.Error(err),
		)
	}
}
