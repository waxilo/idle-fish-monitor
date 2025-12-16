package main

import (
	"test/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", handlers.PingHandler)
	router.Run(":8080")
}
