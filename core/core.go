package core

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"taskm/cache"
	"taskm/enum"
	"taskm/model"
)

func LoadCore(envPath string) {
	//读取.env
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//gin模式
	switch os.Getenv("GIN_MODE") {
	case enum.ModeProduct:
		//开启生产模式,忽视报错
		gin.SetMode(gin.ReleaseMode)
	case enum.ModeDevelop:
		//开发模式,显示报错
		gin.SetMode(gin.DebugMode)
	//case "test":
	//	//测试模式,console打印结果
	//	gin.SetMode(gin.TestMode)
	default:
		log.Fatal("no .env info")
	}
	//拼接数据库连接信息
	dsn := model.GetDsn()
	//创建数据库链接
	model.Database(dsn)
	redisConfig := redis.Options{Addr: os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PWD"), // no password set
		DB:       0,                      // use default DB }
	}
	// 初始化redis缓存客户端
	cache.Init(redisConfig)
}
