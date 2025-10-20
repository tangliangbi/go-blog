package main

import (
	"com.tang.blog/cmd"
	"com.tang.blog/utils/logger"
)

func main() {
	// Log application startup
	logger.Log.Info("Starting Blog Backend Application")

	// 创建服务器实例
	server := cmd.NewServer()

	// 初始化服务器
	err := server.Init()
	if err != nil {
		logger.Log.WithField("error", err.Error()).Fatal("Failed to initialize server")
	}

	// 启动服务器
	err = server.Run()
	if err != nil {
		logger.Log.WithField("error", err.Error()).Fatal("Failed to start server")
	}
}
