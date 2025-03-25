package repository

import (
	"fmt"
	"net"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tarantool/go-tarantool"
)

type TarantoolConfig struct {
    Host     string
    Port     string
    User     string
    Password string
}

// backend/internal/repository/tarantool.go
func NewTarantoolDB(cfg TarantoolConfig) (*tarantool.Connection, error) {
    logrus.Info("Starting Tarantool connection...")
    
    var conn *tarantool.Connection
    var err error
    maxAttempts := 15
    delay := 5 * time.Second
    
    for i := 1; i <= maxAttempts; i++ {
        conn, err = tarantool.Connect(
            net.JoinHostPort(cfg.Host, cfg.Port),
            tarantool.Opts{
                User:          cfg.User,
                Pass:          cfg.Password,
                Timeout:       3 * time.Second,
                Reconnect:    1 * time.Second,
                MaxReconnects: 3,
            },
        )
        
        if err == nil {
            if _, pingErr := conn.Ping(); pingErr == nil {
                logrus.Info("Tarantool connection established")
                return conn, nil
            }
        }
        
        logrus.Warnf("Connection attempt %d/%d failed: %v", i, maxAttempts, err)
        time.Sleep(delay)
    }
    
    return nil, fmt.Errorf("failed to connect after %d attempts: %w", maxAttempts, err)
}