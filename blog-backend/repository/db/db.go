package db

import (
	"fmt"

	"com.tang.blog/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// 使用配置文件中的数据库配置
	cfg := config.AppCfg.Mysql.Default

	host := cfg.DbHost
	user := cfg.UserName
	password := cfg.Password
	dbname := cfg.DbName
	port := cfg.DbPort

	// 构建连接到具体数据库的DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		user, password, host, port, dbname, cfg.Charset)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	fmt.Println("Database connected successfully!")
}
