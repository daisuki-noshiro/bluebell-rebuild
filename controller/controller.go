package controller

import "github.com/gin-gonic/gin"

func PingHandler(c *gin.Context) {
	ResponseSuccess(c, gin.H{
		"message": "pong",
	})
}
