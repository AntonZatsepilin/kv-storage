package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AntonZatsepilin/kv-storage.git/internal/handler"
	"github.com/AntonZatsepilin/kv-storage.git/internal/models"
	"github.com/AntonZatsepilin/kv-storage.git/internal/repository"
	"github.com/AntonZatsepilin/kv-storage.git/internal/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tarantool/go-tarantool/v2"

	_ "github.com/AntonZatsepilin/kv-storage.git/docs"
)

// @title KV Storage API
// @version 1.0
// @description Simple Key-Value storage service with Tarantool backend

// @host localhost:8080
// @BasePath /api
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// if err := godotenv.Load(".env"); err != nil {
	// 	logrus.Fatalf("error loading env variables: %s", err.Error())
	// }

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	tarantoolCfg := repository.TarantoolConfig{
		Addres:     viper.GetString("tarantool.address"),
		User:     os.Getenv("TARANTOOL_USER_PASSWORD"),
		Password: os.Getenv("TARANTOOL_USER_PASSWORD"),
		Timeout: viper.GetInt("tarantool.timeout"),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	var db *tarantool.Connection
	var err error
	db, err = repository.NewTarantoolDB(ctx, tarantoolCfg)
	if err != nil {
		logrus.Fatalf("failed to initialize Tarantool: %s", err.Error())
	}

	defer db.CloseGraceful()

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)
	

	srv := new(models.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("kv-storage Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	logrus.Print("kv-storage Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}



func initConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}