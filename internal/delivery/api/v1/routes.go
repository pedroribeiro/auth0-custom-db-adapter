package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/pedroribeiro/users/internal/delivery/api/v1/users"
	"github.com/pedroribeiro/users/internal/repository"
	"github.com/pedroribeiro/users/internal/usecase"
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
