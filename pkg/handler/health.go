package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ping(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	if err := h.services.Health.Ping(ctx); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "not_ready",
			"db":     "unreachable",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ready",
		"db":     "ok",
	})
}
