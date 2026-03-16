package routes

import (
	"nexora-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(
	router *gin.Engine,
	authMiddleware gin.HandlerFunc,
	userController *controllers.UserController,
) {
	user := router.Group("/api/user")

	user.GET("/profile", authMiddleware, userController.GetUser)
}
