package handlers

import (
	"fmt"
	"go-web/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义校验注册信息的函数
func validateRegistrationData(data map[string]interface{}) error {
	// 示例校验：检查用户名和密码是否存在
	if _, ok := data["username"]; !ok {
		return fmt.Errorf("用户名不能为空")
	}
	if _, ok := data["password"]; !ok {
		return fmt.Errorf("密码不能为空")
	}
	// 可根据实际需求添加更多校验逻辑
	return nil
}

// 定义添加数据到数据库的函数
func addToDatabase(user map[string]interface{}) error {
	models.CreateUser(&models.User{
		Username: user["username"].(string),
		Password: user["password"].(string),
	})
	return nil
}

func PostRegister(c *gin.Context) {
	// 获取前端传入的数据
	var requestData map[string]interface{}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 校验注册信息是否合法
	if err := validateRegistrationData(requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 将合法的注册信息添加到数据库
	if err := addToDatabase(requestData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注册信息添加到数据库失败: " + err.Error()})
		return
	}

	log.Printf("requestData %v", requestData)
	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}
