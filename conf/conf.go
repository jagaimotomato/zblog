package conf

import (
	"github.com/joho/godotenv"
	"os"
	"zblog/cache"
	"zblog/model"
	"zblog/util"
)

//Init 初始化配置项目
func Init() {
	// 本地读取环境变量
	godotenv.Load()

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis()
}
