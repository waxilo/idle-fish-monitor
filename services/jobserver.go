package services

import (
	"context"
	"log"
	"time"
)

// JobServer 任务服务器
type JobServer struct {
	ctx    context.Context
	cancel context.CancelFunc
}

// NewJobServer 创建新的任务服务器
func NewJobServer() *JobServer {
	ctx, cancel := context.WithCancel(context.Background())
	return &JobServer{
		ctx:    ctx,
		cancel: cancel,
	}
}

// Name 返回服务名称
func (j *JobServer) Name() string {
	return "JobServer"
}

// Start 启动任务服务器
func (j *JobServer) Start() error {
	log.Println("JobServer starting...")

	// 启动定时任务
	go j.runJobs()

	log.Println("JobServer started successfully - will run jobs every 30 seconds")
	return nil
}

// Stop 停止任务服务器
func (j *JobServer) Stop() error {
	log.Println("JobServer stopping...")
	j.cancel()
	log.Println("JobServer stopped")
	return nil
}

// runJobs 运行任务
func (j *JobServer) runJobs() {
	ticker := time.NewTicker(10 * time.Second) // 改为10秒执行一次，便于观察
	defer ticker.Stop()

	log.Println("JobServer: Started ticker - executing job every 10 seconds")

	for {
		select {
		case <-j.ctx.Done():
			log.Println("JobServer context cancelled, stopping jobs")
			return
		case <-ticker.C:
			log.Println("JobServer: Ticker triggered - executing job")
			j.executeJob()
		}
	}
}

// executeJob 执行具体任务
func (j *JobServer) executeJob() {
	log.Printf("JobServer: Executing scheduled job at %v", time.Now().Format("2006-01-02 15:04:05"))

	// 在这里添加你的具体任务逻辑
	// 例如：数据同步、清理临时文件、发送通知等

	// 模拟一些工作
	time.Sleep(1 * time.Second)
	log.Println("JobServer: Job completed successfully")
}
