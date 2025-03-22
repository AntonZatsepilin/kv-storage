package repository

import (
	"github.com/sirupsen/logrus"
	"github.com/tarantool/go-tarantool"
)

type TarantoolConfig struct {
    Host     string
    Port     string
    User     string
    Password string
}

func NewTarantoolDB(cfg TarantoolConfig) (*tarantool.Connection, error) {
    logrus.Info("Connecting to Tarantool")
    
    conn, err := tarantool.Connect(cfg.Host+":"+cfg.Port, tarantool.Opts{
        User: cfg.User,
        Pass: cfg.Password,
    })
    
    if err != nil {
        return nil, err
    }
    
	logrus.Info("Checking the connection to the database")
    _, err = conn.Ping()
    if err != nil {
        return nil, err
    }
    
    logrus.Info("Successfully connected to Tarantool")
    return conn, nil
}