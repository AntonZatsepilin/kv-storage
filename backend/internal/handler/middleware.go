package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)
func LoggerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        path := c.Request.URL.Path
        
        c.Next()
        
        entry := logrus.WithFields(logrus.Fields{
            "method":      c.Request.Method,
            "path":        path,
            "status":      c.Writer.Status(),
            "latency":     time.Since(start),
            "client_ip":   c.ClientIP(),
            "user_agent":  c.Request.UserAgent(),
        })
        
        if c.Writer.Status() >= 500 {
            entry.Error("server error")
        } else if c.Writer.Status() >= 400 {
            entry.Warn("client error")
        } else {
            entry.Info("request processed")
        }
    }
}
