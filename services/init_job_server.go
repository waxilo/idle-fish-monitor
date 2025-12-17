package services

import (
	"log"

	"github.com/robfig/cron/v3"
)

// 初始化 JOB服务器
type InitJobServer struct {
	cron *cron.Cron
}

// InitJobServer 创建初始化Job服务器
func NewInitJobServer() *InitJobServer {
	return &InitJobServer{
		cron: cron.New(cron.WithSeconds()),
	}
}

// Name 返回服务名称
func (server *InitJobServer) Name() string {
	return "InitJobServer"
}

// Start 启动Job服务器
func (server *InitJobServer) Start() error {
	_, err := server.cron.AddFunc("*/10 * * * * *", initJob)
	if err != nil {
		log.Printf("Failed to add cron job: %v", err)
		return err
	}
	server.cron.Start()
	return nil
}

// Stop 停止Job服务器
func (server *InitJobServer) Stop() error {
	server.cron.Stop()
	return nil
}

func initJob() {
	log.Printf("1")
}
