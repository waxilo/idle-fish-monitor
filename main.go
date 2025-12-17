package main

import (
	"idle-fish-monitor/bootstrap"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 创建bootstrap实例
	app := bootstrap.NewBootstrap(8080)

	// 初始化服务
	app.Initialize()

	// 设置优雅关闭
	go handleGracefulShutdown(app)

	// 启动所有服务
	if err := app.Start(); err != nil {
		log.Fatal("Failed to start services:", err)
	}

	// 保持程序运行，等待关闭信号
	select {}
}

// handleGracefulShutdown 处理优雅关闭
func handleGracefulShutdown(app *bootstrap.Bootstrap) {
	// 创建信号通道
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 等待信号
	<-quit
	log.Println("Shutdown signal received, stopping services...")

	// 停止所有服务
	if err := app.Stop(); err != nil {
		log.Printf("Error during shutdown: %v", err)
	}

	log.Println("All services stopped gracefully")
}
