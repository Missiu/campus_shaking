package database

import (
	"fmt"
	"github.com/spf13/viper"
)

// InitViper 初始化viper
func InitViper() error {
	// 设置配置文件路径，此路径为其他文件调用时
	viper.SetConfigFile("config.yml")
	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置文件错误: %w", err)
	}
	return nil
}

func InitViperOtherPath() error {
	// 设置配置文件路径，此路径为database文件调用时
	viper.SetConfigFile("../config.yml")
	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置文件错误: %w", err)
	}
	return nil
}
