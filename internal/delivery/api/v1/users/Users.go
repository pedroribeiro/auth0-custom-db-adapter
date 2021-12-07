package users

import (
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/elvenworks/users/internal/delivery/api/middleware"
	"github.com/elvenworks/users/internal/domain/entity"
	"github.com/elvenworks/users/internal/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUseCase usecase.IUserUseCase
}

// InitRoutes config routes
func (c UserHandler) InitRoutes(router *gin.RouterGroup) {
	router.POST("/user", c.Create)
	router.GET("/user/:email", jwt.CheckJWT, c.GetByEmail)
	router.GET("/login", c.Login)
}

// Create godoc
// @Create creates an delivery obj
// @Description creates an delivery obj
// @ID create
// @Accept  json
// @Produce  json
// @Param data body entity.Delivery true "Create deliveryUsecase"
// @Success 201 {object} entity.Delivery
// @Failure 400 {string} http.StatusBadRequest
// @Failure 404 {string} http.StatusNotFound
// @Failure 500 {string} http.StatusInternalServerError
// @Router /delivery [post]
func (c UserHandler) GetByEmail(ctx *gin.Context) {

	email := ctx.Param("email")

	user, err := c.UserUseCase.GetByEmail(email)

	if err != nil {

		if err.Error() == "record not found" {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusAccepted, user)
}

func (c UserHandler) Create(ctx *gin.Context) {
	var user *entity.User

	if err := ctx.Bind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	err := c.UserUseCase.Create(user)

	if err != nil {
		status := http.StatusInternalServerError
		if strings.Contains(err.Error(), "not authorized") {
			status = http.StatusForbidden
		}

		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}

	ctx.String(http.StatusCreated, "created")
}

func (c UserHandler) Login(ctx *gin.Context) {

	username, password, hasAuth := ctx.Request.BasicAuth()

	if !hasAuth {
		ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("missing auth params"))
	}

	user, err := c.UserUseCase.Login(username, password)

	if err != nil {
		status := http.StatusInternalServerError

		if strings.Contains(err.Error(), "wrong password") {
			status = http.StatusForbidden
		}

		if err.Error() == "record not found" {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}

		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
