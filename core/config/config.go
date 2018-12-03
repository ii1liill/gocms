package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// 初始化viper的函数
func SetConfigPath (configPath string) {
	viper.SetConfigName("config")
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

// Get
func Get(key string) interface{} {
	value := viper.Get(key)
	return value
}

//
func GetWithCallback(key string, defaultValue interface{}) interface{} {
	value := viper.Get(key)
	if value == nil {
		return defaultValue
	}
	return value
}