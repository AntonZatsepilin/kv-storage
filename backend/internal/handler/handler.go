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

    router.POST("/register", Handler.Register)
    // router.POST("/login", Handler.Login)
    // router.POST("/refresh", Handler.RefreshTokens)

	return router
}