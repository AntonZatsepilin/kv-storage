package handler

import (
	"github.com/AntonZatsepilin/kv-storage.git/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}


func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(LoggerMiddleware())

	api := router.Group("/api")
	{
		api.POST("/kv", h.setValue)
		api.GET("/kv/:key", h.getValueByKey)
		api.PUT("/kv/:key", h.updateValue)
		api.DELETE("/kv/:key", h.deleteValue)
	}
	return router
}