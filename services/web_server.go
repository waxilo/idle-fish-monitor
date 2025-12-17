package services

import (
	"idle-fish-monitor/services/webserver"

	"github.com/gin-gonic/gin"
)

// WebServer HTTP服务器
type WebServer struct {
	addr   string
	engine *gin.Engine
}

// NewWebServer 创建新的Web服务器
func NewWebServer(addr string) *WebServer {
	return &WebServer{
		addr:   addr,
		engine: gin.Default(),
	}
}

// Name 返回服务名称
func (w *WebServer) Name() string {
	return "WebServer"
}

// Start 启动Web服务器
func (w *WebServer) Start() error {
	w.setupRoutes()
	return w.engine.Run(w.addr)
}

// Stop 停止Web服务器
func (w *WebServer) Stop() error {
	return nil
}

// setupRoutes 设置路由
func (w *WebServer) setupRoutes() {
	w.engine.GET("/ping", webserver.PingHandler)
	w.engine.GET("/search", webserver.SearchHandler)
}
