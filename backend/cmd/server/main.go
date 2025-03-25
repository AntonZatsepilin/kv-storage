package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tarantool/go-tarantool/v2"
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
		if err := godotenv.Load(".env"); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	dialer := tarantool.NetDialer{
		Address:  "tarantool:3301",
		User:     os.Getenv("TARANTOOL_USER_NAME"),
		Password: os.Getenv("TARANTOOL_USER_PASSWORD"),
	}
	opts := tarantool.Opts{
		Timeout: 10 * time.Second,
	}

	conn, err := tarantool.Connect(ctx, dialer, opts)
	if err != nil {
		fmt.Println("Connection refused:", err)
		return
	}

	_, err = conn.Ping()
	if err != nil {
    	logrus.Fatalf("Ping failed: %v", err)
	}

	logrus.Info("Connected to Tarantool")

	// Insert data
tuples := [][]interface{}{
	{1, "Roxette", 1986},
	{2, "Scorpions", 1965},
	{3, "Ace of Base", 1987},
	{4, "The Beatles", 1960},
}
var futures []*tarantool.Future
for _, tuple := range tuples {
	request := tarantool.NewInsertRequest("bands").Tuple(tuple)
	futures = append(futures, conn.Do(request))
}
fmt.Println("Inserted tuples:")
for _, future := range futures {
	result, err := future.Get()
	if err != nil {
		fmt.Println("Got an error:", err)
	} else {
		fmt.Println(result)
	}
}



	// // Select by primary key
	// data, err := conn.Do(
	// 	tarantool.NewSelectRequest("bands").
	// 		Limit(10).
	// 		Iterator(tarantool.IterEq).
	// 		Key([]interface{}{uint(1)}),
	// ).Get()
	// if err != nil {
	// 	fmt.Println("Got an error:", err)
	// }
	// fmt.Println("Tuple selected by the primary key value:", data)

	// // Select by secondary key
	// data, err = conn.Do(
	// 	tarantool.NewSelectRequest("bands").
	// 		Index("band").
	// 		Limit(10).
	// 		Iterator(tarantool.IterEq).
	// 		Key([]interface{}{"The Beatles"}),
	// ).Get()
	// if err != nil {
	// 	fmt.Println("Got an error:", err)
	// }
	// fmt.Println("Tuple selected by the secondary key value:", data)

	// // Update
	// data, err = conn.Do(
	// 	tarantool.NewUpdateRequest("bands").
	// 		Key(tarantool.IntKey{2}).
	// 		Operations(tarantool.NewOperations().Assign(1, "Pink Floyd")),
	// ).Get()
	// if err != nil {
	// 	fmt.Println("Got an error:", err)
	// }
	// fmt.Println("Updated tuple:", data)

	// // Upsert
	// data, err = conn.Do(
	// 	tarantool.NewUpsertRequest("bands").
	// 		Tuple([]interface{}{uint(5), "The Rolling Stones", 1962}).
	// 		Operations(tarantool.NewOperations().Assign(1, "The Doors")),
	// ).Get()
	// if err != nil {
	// 	fmt.Println("Got an error:", err)
	// }

	// // Replace
	// data, err = conn.Do(
	// 	tarantool.NewReplaceRequest("bands").
	// 		Tuple([]interface{}{1, "Queen", 1970}),
	// ).Get()
	// if err != nil {
	// 	fmt.Println("Got an error:", err)
	// }
	// fmt.Println("Replaced tuple:", data)

	// // Delete
	// data, err = conn.Do(
	// 	tarantool.NewDeleteRequest("bands").
	// 		Key([]interface{}{uint(5)}),
	// ).Get()
	// if err != nil {
	// 	fmt.Println("Got an error:", err)
	// }
	// fmt.Println("Deleted tuple:", data)

	// // Call
	// data, err = conn.Do(
	// 	tarantool.NewCallRequest("get_bands_older_than").Args([]interface{}{1966}),
	// ).Get()
	// if err != nil {
	// 	fmt.Println("Got an error:", err)
	// }
	// fmt.Println("Stored procedure result:", data)

	// // Close connection
	// conn.CloseGraceful()
	// fmt.Println("Connection is closed")
}

func initConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}