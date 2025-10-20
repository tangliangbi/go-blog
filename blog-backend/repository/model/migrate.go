package model

import (
	"com.tang.blog/repository/db"
	"com.tang.blog/utils/logger"
	"gorm.io/gorm"
)

// Migrate 执行数据库迁移
func Migrate() {
	// 自动迁移模型
	err := db.GetDB().AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		logger.Log.WithField("error", err.Error()).Fatal("Failed to migrate database")
	}

	logger.Log.Info("Database migration completed!")
}

// MigrateWithDB 使用指定数据库实例执行迁移
func MigrateWithDB(database *gorm.DB) {
	// 自动迁移模型
	err := database.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		logger.Log.WithField("error", err.Error()).Fatal("Failed to migrate database")
	}

	logger.Log.Info("Database migration completed!")
}
