package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health godoc
// @Summary      Health check
// @Description  Returns service health
// @Tags         meta
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /health [get]
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// Version godoc
// @Summary      Version
// @Description  Returns app version
// @Tags         meta
// @Produce      json
// @Success      200 {object} map[string]any
// @Router       /version [get]
func Version(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"app":     "gotpservice",
		"version": "0.1.0-step1",
	})
}
