package handler

import (
	"net/http"

	"github.com/AntonZatsepilin/kv-storage.git/internal/models"
	"github.com/AntonZatsepilin/kv-storage.git/internal/repository"
	"github.com/gin-gonic/gin"
)

// SetValue godoc
// @Summary Create a new key-value pair
// @Description Create new key-value entry in storage
// @Tags kv
// @Accept  json
// @Produce json
// @Param input body models.KeyValue true "Key-Value data"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 409 {object} errorResponse
// @Router /kv [post]
func (h *Handler) setValue(c *gin.Context) {
    var input models.KeyValue

    if err := c.BindJSON(&input); err != nil {
        newErrorResponse(c, http.StatusBadRequest, "invalid JSON format")
        return
    }

    if err := h.services.SetValue(input.Key, string(input.Value)); err != nil {
            newErrorResponse(c, http.StatusConflict, err.Error())
            return
        }

    c.JSON(http.StatusOK, statusResponse{"value set successfully"})
}

// GetValueByKey godoc
// @Summary Get value by key
// @Description Get existing value by key
// @Tags kv
// @Produce json
// @Param key path string true "Key"
// @Success 200 {object} models.KeyValueResp
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /kv/{key} [get]
func (h *Handler) getValueByKey(c *gin.Context) {
    key := c.Param("key")

    value, err := h.services.GetValueByKey(key)
    if err != nil {
        if err == repository.ErrKeyNotFound {
            newErrorResponse(c, http.StatusNotFound, err.Error())
        } else {
            newErrorResponse(c, http.StatusInternalServerError, err.Error())
        }
        return
    }

    c.JSON(http.StatusOK, models.KeyValueResp{Key: key, Value: value})
}

// UpdateValue godoc
// @Summary Update existing value
// @Description Update value for existing key
// @Tags kv
// @Accept  json
// @Produce json
// @Param key path string true "Key"
// @Param input body models.KeyValueUpdateReq true "New value"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Router /kv/{key} [put]
func (h *Handler) updateValue(c *gin.Context) {
	key := c.Param("key")

	var input models.KeyValueUpdateReq

    if err := c.BindJSON(&input); err != nil {
        newErrorResponse(c, http.StatusBadRequest, "invalid JSON format")
        return
    }

	if err := h.services.UpdateValue(key, string(input.Value)); err != nil {
		newErrorResponse(c, http.StatusConflict, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"value updated successfully"})
}

// DeleteValue godoc
// @Summary Delete key-value pair
// @Description Delete existing key-value entry
// @Tags kv
// @Produce json
// @Param key path string true "Key"
// @Success 200 {object} statusResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /kv/{key} [delete]
func (h *Handler) deleteValue(c *gin.Context) {
    key := c.Param("key")

    if err := h.services.DeleteValue(key); err != nil {
        if err == repository.ErrKeyNotFound {
            newErrorResponse(c, http.StatusNotFound, err.Error())
        } else {
            newErrorResponse(c, http.StatusInternalServerError, err.Error())
        }
        return
    }

    c.JSON(http.StatusOK, statusResponse{"value deleted successfully"})
}