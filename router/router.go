package router

import (
	"bluebell_rebuild/controller"

	"github.com/gin-gonic/gin"
)

// SetupRouter 创建并配置 Gin 路由
func SetupRouter() *gin.Engine {
	// 创建一个自带日志和异常恢复功能的 Gin 引擎
	r := gin.Default()

	//注册路由
	r.GET("/ping", controller.PingHandler)
	return r
}
