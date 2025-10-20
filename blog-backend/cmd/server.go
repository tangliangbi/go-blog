package cmd

import (
	"log"

	"com.tang.blog/api/routers"
	"com.tang.blog/config"
	"com.tang.blog/repository/db"
	"com.tang.blog/repository/model"

	"github.com/gin-gonic/gin"
)

// Server 应用服务器结构体
type Server struct {
	config   *config.Config
	router   *gin.Engine
	database *db.Database
}

// NewServer 创建新的服务器实例
func NewServer() *Server {
	return &Server{}
}

// Init 初始化服务器
func (s *Server) Init() error {
	// 初始化配置
	err := config.InitConfig()
	if err != nil {
		return err
	}
	s.config = config.AppCfg

	// 初始化数据库 (方法1: 使用NewDatabase)
	database, err := db.NewDatabase(&s.config.Mysql)
	if err != nil {
		// 如果方法1失败，尝试方法2: 使用InitDB
		db.InitDB()
	} else {
		s.database = database
		// 设置全局DB变量
		db.DB = database.DB
	}

	// 执行数据库迁移
	if db.GetDB() != nil {
		model.MigrateWithDB(db.GetDB())
	}

	// 初始化路由
	s.initRouter()

	return nil
}

// initRouter 初始化路由
func (s *Server) initRouter() {
	s.router = gin.Default()
	routers.SetupRoutes(s.router)
}

// Run 启动服务器
func (s *Server) Run() error {
	addr := s.config.System.Host + s.config.System.HttpPort
	log.Println("Server starting on", addr)
	return s.router.Run(addr)
}
