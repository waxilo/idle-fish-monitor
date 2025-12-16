package handlers

import (
	"github.com/gin-gonic/gin"
)

// PingHandler 处理 /ping 请求
func PingHandler(response *gin.Context) {
	response.JSON(200, gin.H{
		"message": "pong",
	})
}
