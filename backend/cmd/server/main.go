package main

import (
	"github.com/AntonZatsepilin/kv-storage.git/internal/handler"
	"github.com/AntonZatsepilin/kv-storage.git/internal/models"
	"github.com/AntonZatsepilin/kv-storage.git/internal/repository"
	"github.com/AntonZatsepilin/kv-storage.git/internal/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	tarantoolCfg := repository.TarantoolConfig{
		Host:     viper.GetString("tarantool.host"),
		Port:     viper.GetString("tarantool.port"),
		User:     viper.GetString("tarantool.user"),
		Password: viper.GetString("tarantool.password"),
	}

	conn, err := repository.NewTarantoolDB(tarantoolCfg)
	if err != nil {
		logrus.Fatalf("failed to initialize Tarantool: %s", err.Error())
	}

	repo := repository.NewRepository(conn)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(models.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occurred while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}