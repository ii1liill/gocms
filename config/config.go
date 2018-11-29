package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func SetConfigPath (configPath string) {
	viper.SetConfigName("config")
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func Get(key string) interface{} {
	return viper.Get(key)
}
