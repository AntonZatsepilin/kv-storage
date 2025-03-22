package handler

import (
	"net/http"

	"github.com/AntonZatsepilin/kv-storage.git/internal/models"
	"github.com/AntonZatsepilin/kv-storage.git/internal/service"
	"github.com/gin-gonic/gin"
)

func (h *Handler) create(c *gin.Context) {
	var req models.KVRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	
	if err := h.services.Create(req.Key, req.Value); err != nil {
		if err == service.ErrKeyExists {
			newErrorResponse(c, http.StatusConflict, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, "internal error")
		}
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{"status": "created"})
}

func (h *Handler) get(c *gin.Context) {
	key := c.Param("id")
	
	value, err := h.services.Get(key)
	if err != nil {
		if err == service.ErrKeyNotFound {
			newErrorResponse(c, http.StatusNotFound, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, "internal error")
		}
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"value": value})
}

func (h *Handler) update(c *gin.Context) {
	key := c.Param("id")
	var req struct {
		Value interface{} `json:"value"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	
	if err := h.services.Update(key, req.Value); err != nil {
		if err == service.ErrKeyNotFound {
			newErrorResponse(c, http.StatusNotFound, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, "internal error")
		}
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *Handler) delete(c *gin.Context) {
	key := c.Param("id")
	
	if err := h.services.Delete(key); err != nil {
		if err == service.ErrKeyNotFound {
			newErrorResponse(c, http.StatusNotFound, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, "internal error")
		}
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}