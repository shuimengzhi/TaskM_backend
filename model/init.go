package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

// DB 数据库链接单例
var DB *gorm.DB

// Database 在中间件中初始化mysql链接
func Database(connString string) {
	// 初始化GORM日志配置
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             20 * time.Second, // Slow SQL threshold
			LogLevel:                  logger.Warn,      // Log level(这里记得根据需求改一下)
			IgnoreRecordNotFoundError: true,             // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,             // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "tm_", // table name prefix, table for `User` would be `t_users`
		},
	})
	// Error
	if connString == "" || err != nil {
		fmt.Printf("mysql lost: %v", err)
		panic(err)
	}
	//sqlDB, err := db.DB()
	//if err != nil {
	//	fmt.Printf("mysql lost: %v", err)
	//	panic(err)
	//}

	//设置连接池
	//空闲
	//sqlDB.SetMaxIdleConns(10)
	////打开
	//sqlDB.SetMaxOpenConns(20)
	DB = db

}

// GetDsn 获得.env配置信息组合成数据库链接
func GetDsn() string {
	envSlice := []string{"MYSQL_USER", "MYSQL_PASSWORD", "MYSQL_HOST", "MYSQL_PORT", "MYSQL_DB_NAME"}
	for _, value := range envSlice {
		if os.Getenv(value) == "" {
			panic(value + "环境变量不存在")
		}
	}
	//"user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASSWORD") +
		"@tcp(" + os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT") +
		")/" + os.Getenv("MYSQL_DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	return dsn
}

// Generator 根据数据库的表生成struct
func Generator(tableName string) {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(DB) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	//g.ApplyBasic(model.User{})
	g.GenerateModel(tableName)
	// Generate the code
	g.Execute()
}
