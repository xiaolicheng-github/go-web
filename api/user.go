package api

import (
	"go-web/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserRouter 初始化用户相关的路由
func UserRouter(r *gin.RouterGroup) {
	userGroup := r.Group("/users")
	userGroup.Use(middleware.AuthMiddleware())
	{
		userGroup.GET("", GetUsers)
		userGroup.GET("/:id", GetUserByID)
		userGroup.POST("", CreateUser)
		userGroup.PUT("/:id", UpdateUser)
		userGroup.DELETE("/:id", DeleteUser)
	}
}

// GetUsers 获取用户列表
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "获取用户列表"})
}

// GetUserByID 根据 ID 获取用户
func GetUserByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "根据 ID 获取用户"})
}

// CreateUser 创建新用户
func CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "创建新用户"})
}

// UpdateUser 更新用户信息
func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "更新用户信息"})
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "删除用户"})
}
