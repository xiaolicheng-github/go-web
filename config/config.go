package config

import (
	"log"
	"os"
)

// GetGinMode 获取 GIN_MODE 环境变量的值
func GetGinMode() string {
	log.Println("当前环境:", os.Getenv("GIN_MODE"))
	return os.Getenv("GIN_MODE")
}

func IsRelease() bool {
	return GetGinMode() == "release"
}

func IsDebug() bool {
	return GetGinMode() == "debug"
}
