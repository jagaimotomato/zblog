package model

import (
	"log"
	"os"
	"time"
	"zblog/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB 数据库链接单例
var DB *gorm.DB

// Database 在中间件中初始化mysql链接
func Database(connString string) {
	// 初始化GORM日志配置
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			Colorful:                  false,       // Disable color
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			LogLevel:                  logger.Info,
		},
	)

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		Logger: newLogger,
	})
	// Error
	if connString == "" || err != nil {
		util.Log().Error("mysql lost: %v", err)
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		util.Log().Error("mysql lost: %v", err)
		panic(err)
	}

	// 设置连接池
	// 空闲
	sqlDB.SetMaxIdleConns(10)
	// 打开
	sqlDB.SetMaxOpenConns(20)
	DB = db

	migration()
}