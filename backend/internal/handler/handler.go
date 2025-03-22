package handler

import (
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

	router.POST("/kv", h.createUser)
	// router.PUT("/kv/:id", h.updateUser)
	// router.GET("/kv/:id", h.getUser)
	// router.DELETE("/kv/:id", h.deleteUser)

	return router
}