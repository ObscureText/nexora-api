package controllers

import (
	"net/http"
	"nexora-api/internal/constants"
	"nexora-api/internal/controllers/interceptor"
	"nexora-api/internal/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (userController *UserController) GetUser(ctx *gin.Context) {
	userId := ctx.MustGet(constants.KEY_USER_ID).(string)

	user, getErr := userController.userService.GetUser(userId)
	if getErr != nil {
		interceptor.HandleServiceError(ctx, getErr)
		return
	}

	ctx.JSON(http.StatusOK, user)
}
