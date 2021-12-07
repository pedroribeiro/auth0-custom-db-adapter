package v1

import (
	"github.com/elvenworks/users/internal/delivery/api/v1/users"
	"github.com/elvenworks/users/internal/repository"
	"github.com/elvenworks/users/internal/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(router *gin.RouterGroup, db *gorm.DB) {

	userRepo := repository.UserRepo{DB: db}

	userUsecase := usecase.UserUseCase{
		Repo: &userRepo,
	}

	userHandler := users.UserHandler{
		UserUseCase: &userUsecase,
	}

	userHandler.InitRoutes(router)

}
