package routes

import (
	"nexora-api/internal/controllers"
	"nexora-api/internal/domain"
	"nexora-api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(
	router *gin.Engine,
	authMiddleware gin.HandlerFunc,
	userController *controllers.UserController,
) {
	user := router.Group("/api/")

	user.GET("/admin/profile", authMiddleware, middlewares.RequireRoles(domain.ServiceProviderAdmin), userController.GetUser)
	user.GET("/individual/profile", authMiddleware, middlewares.RequireRoles(domain.Individual), userController.GetUser)
	user.GET("/establishment/profile", authMiddleware, middlewares.RequireRoles(domain.EstablishmentOwner), userController.GetUser)
}
