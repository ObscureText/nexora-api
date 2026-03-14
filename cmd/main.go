package main

import (
	"log"
	"nexora-api/internal/controllers"
	"nexora-api/internal/db"
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
	adminRepository := repositories.NewAdminRepository(database)

	authService := services.NewAuthService(
		jwtUtil,
		cryptUtil,
		userRepository,
		adminRepository,
	)

	authController := controllers.NewAuthController(authService)

	router := gin.Default()

	routes.RegisterAuthRoutes(router, authController)

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Nexora is alive ...")
	})

	router.Run(":8080")
}
