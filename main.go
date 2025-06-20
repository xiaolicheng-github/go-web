package main

import (
	"go-web/handlers"
	"go-web/middleware"
	"go-web/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	// 连接SQLite数据库
	db, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}

	// 自动迁移模型
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic("数据库迁移失败: " + err.Error())
	}
	models.DB = db // 赋值全局DB实例

	// 初始化Gin路由
	router := gin.Default()

	// 提供首页静态文件
	router.StaticFile("/", "./web/index.html")

	// 连接数据库公共路由组（无需认证）
	publicGroup := router.Group("/api")
	{
		publicGroup.POST("/register", handlers.Register)
		publicGroup.POST("/login", handlers.Login)
	}

	// 认证路由组（需要JWT）
	authGroup := router.Group("/api")
	authGroup.Use(middleware.AuthMiddleware())
	{
		authGroup.GET("/user", handlers.GetUser)
		authGroup.PUT("/user", handlers.UpdateUser)
		authGroup.DELETE("/user", handlers.DeleteUser)
	}

	// 启动服务
	router.Run(":4000")
}
