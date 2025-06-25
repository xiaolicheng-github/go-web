package main

import (
	"go-web/handlers"
	"go-web/middleware"
	"go-web/models"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

var (
	g errgroup.Group
)

func webRouter() http.Handler {

	if os.Getenv("GIN_MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()

	router.Static("/", "./web/dist")

	return router

}
func apiRouter() http.Handler {
	db, err := gorm.Open(sqlite.Open("go-web.db"), &gorm.Config{})
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
	// 生产环境关闭调试模式
	if os.Getenv("GIN_MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
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

	return router
}

func main() {
	webServer := &http.Server{
		Addr:    ":4000",
		Handler: webRouter(),
	}
	apiServer := &http.Server{
		Addr:    ":4001",
		Handler: apiRouter(),
	}

	g.Go(func() error {
		return webServer.ListenAndServe()
	})

	g.Go(func() error {
		return apiServer.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
