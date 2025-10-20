package db

import (
	"fmt"

	"com.tang.blog/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database 数据库结构体
type Database struct {
	*gorm.DB
}

// NewDatabase 创建新的数据库实例
func NewDatabase(cfg *config.MysqlConfig) (*Database, error) {
	host := cfg.Default.DbHost
	user := cfg.Default.UserName
	password := cfg.Default.Password
	dbname := cfg.Default.DbName
	port := cfg.Default.DbPort

	// 构建连接到具体数据库的DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		user, password, host, port, dbname, cfg.Default.Charset)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Database connected successfully!")

	return &Database{DB: db}, nil
}

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	return DB
}
