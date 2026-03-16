package main

import (
	"log"
	"nexora-api/internal/controllers"
	"nexora-api/internal/db"
	middleware "nexora-api/internal/middlewares"
	"nexora-api/internal/repositories"
	"nexora-api/internal/routes"
	"nexora-api/internal/services"
	"nexora-api/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatal("Database connection failed : ", err)
	}

	jwtUtil := utils.NewJWTUtil()
	cryptUtil := utils.NewCryptoUtils()

	userRepository := repositories.NewUserRepository(database)

	userService := services.NewUserService(userRepository)
	authService := services.NewAuthService(
		jwtUtil,
		cryptUtil,
		userRepository,
	)

	userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(authService)

	router := gin.Default()

	authMiddleware := middleware.AuthMiddleware(jwtUtil)

	routes.RegisterUserRoutes(router, authMiddleware, userController)
	routes.RegisterAuthRoutes(router, authController)

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Nexora is live ...")
	})

	router.Run(":8080")
}
