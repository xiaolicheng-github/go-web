package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"go-web/models"
	"time"
)

// Register 用户注册
func Register(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "请求参数错误"})
		return
	}

	// 检查用户名是否已存在
	var existing models.User
	if err := models.DB.Where("username = ?", req.Username).First(&existing).Error; err == nil {
		c.JSON(400, gin.H{"error": "用户名已存在"})
		return
	}

	// 创建新用户
	user := models.User{Username: req.Username, Password: req.Password}
	if err := models.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "注册失败"})
		return
	}

	c.JSON(200, gin.H{"message": "注册成功"})
}

// Login 用户登录
func Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "请求参数错误"})
		return
	}

	// 查询用户
	var user models.User
	if err := models.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(400, gin.H{"error": "用户不存在"})
		} else {
			c.JSON(500, gin.H{"error": "登录失败"})
		}
		return
	}

	// 验证密码
	if !user.CheckPassword(req.Password) {
		c.JSON(400, gin.H{"error": "密码错误"})
		return
	}

	// 生成JWT令牌
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("your-secret-key")) // 建议从配置文件读取密钥
	if err != nil {
		c.JSON(500, gin.H{"error": "生成令牌失败"})
		return
	}

	c.JSON(200, gin.H{"token": tokenString})
}

// GetUser 获取当前用户信息
func GetUser(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var user models.User
	if err := models.DB.First(&user, userID).Error; err != nil {
		c.JSON(500, gin.H{"error": "获取用户信息失败"})
		return
	}

	// 过滤敏感字段
	user.Password = ""
	c.JSON(200, user)
}

// UpdateUser 更新用户信息
func UpdateUser(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var req struct {
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "请求参数错误"})
		return
	}

	var user models.User
	if err := models.DB.First(&user, userID).Error; err != nil {
		c.JSON(500, gin.H{"error": "用户不存在"})
		return
	}

	user.Password = req.Password
	if err := models.DB.Save(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "更新失败"})
		return
	}

	c.JSON(200, gin.H{"message": "更新成功"})
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	if err := models.DB.Delete(&models.User{}, userID).Error; err != nil {
		c.JSON(500, gin.H{"error": "删除失败"})
		return
	}

	c.JSON(200, gin.H{"message": "删除成功"})
}