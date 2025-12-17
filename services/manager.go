package services

import (
	"log"
	"sync"
)

// Manager 服务管理器
type Manager struct {
	services []Service
	mu       sync.RWMutex
}

// NewManager 创建新的服务管理器
func NewManager() *Manager {
	return &Manager{
		services: make([]Service, 0),
	}
}

// Register 注册服务
func (m *Manager) Register(service Service) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.services = append(m.services, service)
	log.Printf("Service registered: %s", service.Name())
}

// StartAll 启动所有服务
func (m *Manager) StartAll() error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, service := range m.services {
		log.Printf("Starting service: %s", service.Name())
		if err := service.Start(); err != nil {
			log.Printf("Failed to start service %s: %v", service.Name(), err)
			return err
		}
		log.Printf("Service started successfully: %s", service.Name())
	}
	return nil
}

// StopAll 停止所有服务
func (m *Manager) StopAll() error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	// 逆序停止服务
	for i := len(m.services) - 1; i >= 0; i-- {
		service := m.services[i]
		log.Printf("Stopping service: %s", service.Name())
		if err := service.Stop(); err != nil {
			log.Printf("Failed to stop service %s: %v", service.Name(), err)
			return err
		}
		log.Printf("Service stopped successfully: %s", service.Name())
	}
	return nil
}

// GetService 根据名称获取服务
func (m *Manager) GetService(name string) Service {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, service := range m.services {
		if service.Name() == name {
			return service
		}
	}
	return nil
}
