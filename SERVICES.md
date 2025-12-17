# 服务封装架构

## 项目结构

```
.
├── main.go                 # 应用入口
├── bootstrap/              # 引导包
│   └── bootstrap.go        # 服务启动引导
├── services/               # 服务包
│   ├── interface.go        # 服务接口定义
│   ├── manager.go          # 服务管理器
│   ├── webserver.go        # Web服务器实现
│   └── jobserver.go        # 任务服务器实现
├── handlers/               # HTTP处理器
│   └── PingHandler.go      # Ping处理器
└── servers/                # 旧的服务文件（可删除）
```

## 核心组件

### 1. Service 接口
```go
type Service interface {
    Start() error    // 启动服务
    Stop() error     // 停止服务
    Name() string    // 服务名称
}
```

### 2. Manager 服务管理器
- 管理所有服务的生命周期
- 支持注册、启动、停止服务
- 提供服务查询功能

### 3. Bootstrap 引导器
- 负责服务的初始化和启动
- 处理优雅关闭

## 使用方法

### 1. 创建新服务
实现 `Service` 接口：

```go
type MyService struct {}

func (m *MyService) Name() string {
    return "MyService"
}

func (m *MyService) Start() error {
    // 启动逻辑
    return nil
}

func (m *MyService) Stop() error {
    // 停止逻辑
    return nil
}
```

### 2. 注册服务
在 `bootstrap/bootstrap.go` 中注册：

```go
func (b *Bootstrap) Initialize() {
    // 注册现有服务
    webServer := services.NewWebServer(":8080")
    b.manager.Register(webServer)
    
    // 注册新服务
    myService := &MyService{}
    b.manager.Register(myService)
}
```

### 3. 运行应用
```bash
go run main.go
```

## 特性

- **统一管理**: 所有服务通过统一接口管理
- **优雅关闭**: 支持信号处理和优雅关闭
- **扩展性**: 易于添加新的服务类型
- **并发安全**: 服务管理器使用读写锁保证线程安全
- **日志记录**: 完整的服务生命周期日志

## 信号处理

应用支持以下信号：
- `SIGINT` (Ctrl+C)
- `SIGTERM` (终止信号)

收到信号后会优雅关闭所有服务。