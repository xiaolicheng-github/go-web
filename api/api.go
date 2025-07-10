package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiRouter(r *gin.Engine) {
	r.Use(func(c *gin.Context) {
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
	apiGroup := r.Group("/api")
	{
		UserRouter(apiGroup)
		LoginRouter(apiGroup)
	}
}
