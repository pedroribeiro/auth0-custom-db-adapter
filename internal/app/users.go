package app

import (
	"os"

	"github.com/elvenworks/users/internal/delivery/api"
	migration "github.com/elvenworks/users/internal/domain/migrations"
	"github.com/elvenworks/users/internal/driver/database"
	"github.com/elvenworks/users/internal/driver/logs"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Start(router *gin.Engine) {

	postgresConn := os.Getenv("DB_HOST")

	logs.Init()

	db := database.New(postgresConn)

	migration.Setup(db)

	logrus.Info("Initializing Workers...")
	api.InitRoutes(router, db)
}
