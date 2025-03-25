package handler

import (
	"net/http"

	"github.com/AntonZatsepilin/kv-storage.git/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) setValue(c *gin.Context) {
	var input models.KeyValue
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.SetValue(input.Key, input.Value); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"value set successfully"})
}

func (h *Handler) getValueByKey(c *gin.Context) {
	key := c.Param("key")

	value, err := h.services.GetValueByKey(key)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.KeyValue{Key: key, Value: value})
}