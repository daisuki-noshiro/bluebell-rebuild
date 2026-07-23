package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    1000,
		Message: "success",
		Data:    data,
	})
}
