package api

import (
	"go-web/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginRouter(r *gin.RouterGroup) {
	loginGroup := r.Group("/login")
	{
		loginGroup.POST("", PostLogin)
		loginGroup.POST("/register", handlers.PostRegister)
	}
}

func PostLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "获取用户登录信息"})
}
