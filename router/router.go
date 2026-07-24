package router

import (
	"bluebell_rebuild/controller"
	"bluebell_rebuild/logger"

	"github.com/gin-gonic/gin"
)

// SetupRouter 创建并配置Gin路由
func SetupRouter() *gin.Engine {
	// 创建一个不自带中间件的Gin引擎
	r := gin.New()

	// 使用Zap请求日志和Gin异常恢复
	r.Use(
		logger.GinLogger(),
		logger.GinRecovery(true),
	)

	// 注册路由
	r.GET("/ping", controller.PingHandler)

	return r
}
