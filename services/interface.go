package services

// Service 定义服务接口
type Service interface {
	// Start 启动服务
	Start() error
	// Stop 停止服务
	Stop() error
	// Name 返回服务名称
	Name() string
}
