package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthHandler Handler
type HealthHandler struct{}

// InitRoutes init
func (c HealthHandler) InitRoutes(router *gin.Engine) {
	router.GET("/", c.Live)
	router.GET("/healthz", c.Live)
}

// Live send status code 200
func (c HealthHandler) Live(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Running"})
}

// NewHealthHandler Create a new handler
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// InitHealthHandler Auto creator for dependency injection
func InitHealthHandler() *HealthHandler {
	return NewHealthHandler()
}
