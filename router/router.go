package router

import (
	"context"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	"taskm/cache"
	common_controller "taskm/controllers/common"
	project_controller "taskm/controllers/project"
	user_controller "taskm/controllers/user"
	"taskm/enum"
	"taskm/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	// 设置跨域
	r.Use(middleware.Cors())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get pong",
		})
	})
	r.POST("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "post pong",
		})
	})
	r.GET("/redis", func(c *gin.Context) {
		cxt := context.Background()
		result, err := cache.Instance.Get(cxt, "test").Result()
		c.JSON(200, gin.H{
			"message": result,
			"err":     err,
		})
	})

	// 只有开发环境才能够看到路由
	if os.Getenv("GIN_MODE") == enum.ModeDevelop {
		r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	static := r.Group("/static")
	{
		static.Static("/avatar", "./upload/avatar")
	}
	upload := r.Group("/upload")
	{
		uploadAuth := upload.Use(middleware.AuthRequired())
		{
			uploadAuth.POST("/avatar", common_controller.UploadAvatar)
		}

	}
	user := r.Group("/user")
	{
		user.POST("/login", user_controller.Login)
		adminAuth := user.Use(middleware.AdminAuth())
		{
			adminAuth.POST("/register", user_controller.Register)
		}

	}
	project := r.Group("/project")
	{
		project.Use(middleware.AuthRequired())

		project.POST("/create", project_controller.ProjectCreate)

	}
	return r
}
