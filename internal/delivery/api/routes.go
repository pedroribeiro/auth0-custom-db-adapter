package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pedroribeiro/users/internal/delivery/api/handler"
	"github.com/pedroribeiro/users/internal/delivery/api/middleware"
	v1 "github.com/pedroribeiro/users/internal/delivery/api/v1"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"gorm.io/gorm"
)

func InitRoutes(router *gin.Engine, db *gorm.DB) {

	p := ginprometheus.NewPrometheus("gin")
	p.Use(router)

	// Logger
	skipPaths := []string{
		"/healthz",
		"/metrics",
	}
	logger := middleware.NewLogger(true, skipPaths)
	logger.UseLogger(router)

	// middleware.UseCors(router)

	// Handlers
	health := handler.InitHealthHandler()
	health.InitRoutes(router)

	// Routes v1
	v1Group := router.Group("external/v1")
	v1.InitRoutes(v1Group, db)
}
