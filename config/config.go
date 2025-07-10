package config

import (
	"os"
)

// GetGinMode 获取 GIN_MODE 环境变量的值
func GetGinMode() string {
	return os.Getenv("GIN_MODE")
}

func IsRelease() bool {
	return GetGinMode() == "release"
}

func IsDebug() bool {
	return GetGinMode() == "debug"
}
