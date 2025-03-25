package main

import (
	"fmt"
	"log"

	"github.com/tarantool/go-tarantool"
)

func main() {
	// if err := initConfig(); err != nil {
	// 	logrus.Fatalf("error initializing config: %s", err.Error())
	// }

	// tarantoolCfg := repository.TarantoolConfig{
	// 	Host:     viper.GetString("tarantool.host"),
	// 	Port:     viper.GetString("tarantool.port"),
	// 	User:     viper.GetString("tarantool.user"),
	// 	Password: viper.GetString("tarantool.password"),
	// }

	// conn, err := repository.NewTarantoolDB(tarantoolCfg)
	// if err != nil {
	// 	logrus.Fatalf("failed to initialize Tarantool: %s", err.Error())
	// }

	// repo := repository.NewRepository(conn)
	// services := service.NewService(repo)
	// handlers := handler.NewHandler(services)

	// srv := new(models.Server)
	// if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
	// 	logrus.Fatalf("error occurred while running http server: %s", err.Error())
	// }

	opts := tarantool.Opts{
		User: "guest",
		// Password: "your_password", // если используется другой пользователь с паролем
	}

	// Устанавливаем соединение с Tarantool
	conn, err := tarantool.Connect("127.0.0.1:3301", opts)
	if err != nil {
		log.Fatalf("Ошибка подключения: %s", err.Error())
	}
	defer conn.Close()

	// Пробуем выполнить команду Ping
	resp, err := conn.Ping()
	if err != nil {
		log.Fatalf("Ошибка при выполнении Ping: %s", err.Error())
	}
	fmt.Println("Ответ на Ping:", resp)
}

// func initConfig() error {
// 	viper.AddConfigPath("./configs")
// 	viper.SetConfigName("config")
// 	return viper.ReadInConfig()
// }