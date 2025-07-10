package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginRouter(r *gin.RouterGroup) {
	loginGroup := r.Group("/login")
	{
		loginGroup.POST("", PostLogin)
		loginGroup.POST("/register", PostRegister)
	}
}

func PostLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "获取用户登录信息"})
}

func PostRegister(c *gin.Context) {
	// 获取前端传入的数据
	var requestData map[string]interface{}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("requestData %v", requestData)
	c.JSON(http.StatusOK, gin.H{"message": "获取用户注册信息"})
}
