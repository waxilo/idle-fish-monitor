package bootstrap

import (
	"idle-fish-monitor/services"
	"log"
	"strconv"
)

// Bootstrap 应用启动引导
type Bootstrap struct {
	manager *services.Manager
	port    int
}

// NewBootstrap 创建新的Bootstrap实例
func NewBootstrap(port int) *Bootstrap {
	return &Bootstrap{
		manager: services.NewManager(),
		port:    port,
	}
}

// Initialize 初始化所有服务
func (b *Bootstrap) Initialize() {
	// 创建Web服务器
	webServer := services.NewWebServer(":" + strconv.Itoa(b.port))
	b.manager.Register(webServer)

	// 创建Job服务器
	jobServer := services.NewJobServer()
	b.manager.Register(jobServer)

	log.Println("All services initialized")
}

// Start 启动所有服务
func (b *Bootstrap) Start() error {
	log.Println("Starting all services...")
	return b.manager.StartAll()
}

// Stop 停止所有服务
func (b *Bootstrap) Stop() error {
	log.Println("Stopping all services...")
	return b.manager.StopAll()
}

// GetManager 获取服务管理器
func (b *Bootstrap) GetManager() *services.Manager {
	return b.manager
}
