package controllers

import (
	"net/http"

	"nexora-api/internal/controllers/dto"
	"nexora-api/internal/controllers/interceptor"

	"nexora-api/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService      services.AuthService
	errorInterceptor interceptor.ErrorInterceptor
}

func NewAuthController(
	authService services.AuthService,
	errorInterceptor interceptor.ErrorInterceptor,
) *AuthController {
	return &AuthController{
		authService:      authService,
		errorInterceptor: errorInterceptor,
	}
}

func (authController *AuthController) Login(context *gin.Context) {
	var loginRequest dto.LoginRequest

	if bindErr := context.ShouldBindJSON(&loginRequest); bindErr != nil {
		authController.errorInterceptor.HandleBadRequest(context, bindErr)
		return
	}

	loginResponse, loginErr := authController.authService.Login(&loginRequest)
	if loginErr != nil {
		authController.errorInterceptor.HandleServiceError(context, loginErr)
		return
	}

	context.JSON(http.StatusOK, loginResponse)
}
