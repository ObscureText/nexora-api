package routes

import (
	"nexora-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(
	router *gin.Engine,
	authController *controllers.AuthController,
) {
	auth := router.Group("/api/auth")

	auth.POST("/login", authController.Login)
}
