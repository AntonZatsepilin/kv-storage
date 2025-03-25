package repository

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tarantool/go-tarantool/v2"
)

type TarantoolConfig struct {
    Addres string
    User string
    Password string
    Timeout int
}

func NewTarantoolDB(ctx context.Context, cfg TarantoolConfig) (*tarantool.Connection, error) {
    logrus.Info("Starting Tarantool connection...")

    dialer := tarantool.NetDialer{
        Address:  cfg.Addres,
        User:     cfg.User,  
        Password: cfg.Password,
    }

    opts := tarantool.Opts{
        Timeout: time.Duration(cfg.Timeout) * time.Second,
    }

    conn, err := tarantool.Connect(ctx, dialer, opts)
    if err != nil {
        logrus.Errorf("Failed to connect to Tarantool: %v", err)
        return nil, err
    }

    if _, err := conn.Ping(); err != nil {
        logrus.Errorf("Ping to Tarantool failed: %v", err)
        return nil, err
    }

    logrus.Info("Successfully connected to Tarantool")
    return conn, nil
}