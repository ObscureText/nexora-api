package controllers

import (
	"net/http"

	"nexora-api/internal/controllers/dto"
	"nexora-api/internal/controllers/interceptor"

	nexoraError "nexora-api/internal/errors"
	"nexora-api/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (authController *AuthController) Login(context *gin.Context) {
	var loginRequest dto.LoginRequest

	if bindErr := context.ShouldBindJSON(&loginRequest); bindErr != nil {
		interceptor.HandleBadRequest(context, bindErr)
		return
	}

	if !loginRequest.UserRole.IsValid() {
		context.JSON(http.StatusBadRequest, nexoraError.BadRequestNexoraError)
		return
	}

	loginResponse, loginErr := authController.authService.Login(&loginRequest)
	if loginErr != nil {
		interceptor.HandleServiceError(context, loginErr)
		return
	}

	context.JSON(http.StatusOK, loginResponse)
}

// router for auth
// db conn file
