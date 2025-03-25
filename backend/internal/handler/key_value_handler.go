package handler

import (
	"encoding/json"
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

	if !json.Valid(input.Value) {
        newErrorResponse(c, http.StatusBadRequest, "value must be a valid JSON")
        return
    }

	if err := h.services.SetValue(input.Key, string(input.Value)); err != nil {
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

	c.JSON(http.StatusOK, models.KeyValueResp{Key: key, Value: value})
}

func (h *Handler) updateValue(c *gin.Context) {
	key := c.Param("key")

	var input models.KeyValue

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if !json.Valid(input.Value) {
        newErrorResponse(c, http.StatusBadRequest, "value must be a valid JSON")
        return
    }

	if err := h.services.UpdateValue(key, string(input.Value)); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"value updated successfully"})
}