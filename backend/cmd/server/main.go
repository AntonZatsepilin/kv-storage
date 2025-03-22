package main

import (
	"github.com/spf13/viper"
)

func main() {

}

func initConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}