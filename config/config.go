package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config/")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Load Config Error: %s", err.Error()))
	}
	fmt.Println("Configuration file loaded successfully")
}
