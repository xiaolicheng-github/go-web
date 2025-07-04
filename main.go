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

	// 处理所有未匹配的路由，指向静态目录的index.html
	router.NoRoute(func(c *gin.Context) {
		c.File("./web/dist/index.html")
	})

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
	// 允许 localhost:4000 进行跨域请求的中间件
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// // 记录请求方法
		// log.Printf("请求方法: %s", c.Request.Method)
		// // 记录请求路径
		// log.Printf("请求路径: %s", c.Request.URL.Path)
		// // 记录客户端IP
		// log.Printf("客户端IP: %s", c.ClientIP())
		// // 记录请求头
		// log.Printf("请求头: %v", c.Request.Header)

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})
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
		authGroup.POST("/user/update", handlers.UpdateUser)
		authGroup.POST("/user/delete", handlers.DeleteUser)
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
