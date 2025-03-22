package handler

import (
	"net/http"

	"github.com/AntonZatsepilin/kv-storage.git/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createUser(c *gin.Context) {
	var input models.KVRequest

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.CreateUser(c.Request.Context(), input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}