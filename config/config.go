package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config/")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	fmt.Println("Configuration file loaded successfully")
	return nil
}
