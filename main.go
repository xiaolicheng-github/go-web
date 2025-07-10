package main

import (
	"go-web/api"
	"go-web/config"
	"go-web/models"
	"log"
	"net/http"

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

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Static("/", "./bin/web")

	// 处理所有未匹配的路由，指向静态目录的index.html
	router.NoRoute(func(c *gin.Context) {
		c.File("./bin/web/index.html")
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

	if config.IsRelease() {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()

	api.ApiRouter(router)

	return router
}

func main() {
	var webServer *http.Server
	if config.IsRelease() {
		webServer = &http.Server{
			Addr:    ":4000",
			Handler: webRouter(),
		}
	}

	apiServer := &http.Server{
		Addr:    ":4001",
		Handler: apiRouter(),
	}
	if config.IsRelease() {
		g.Go(func() error {
			return webServer.ListenAndServe()
		})
	}

	g.Go(func() error {
		return apiServer.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
