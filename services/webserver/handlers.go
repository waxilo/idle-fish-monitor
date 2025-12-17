package webserver

import (
	"github.com/gin-gonic/gin"
)

// PingHandler 处理 /ping 请求
func PingHandler(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong",
	})
}

// SearchHandler 处理商品查询
func SearchHandler(context *gin.Context) {
	// todo 获取参数

	// todo 调用接口

	// todo 返回结果
}
