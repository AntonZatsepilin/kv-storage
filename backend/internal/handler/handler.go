package handler

import (
	"github.com/AntonZatsepilin/kv-storage.git/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}


func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/kv", h.setValue)
		api.GET("/kv/:key", h.getValueByKey)
		api.PUT("/kv/:key", h.updateValue)
	}
	return router
}