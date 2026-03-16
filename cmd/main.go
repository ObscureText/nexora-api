package main

import (
	"log"
	"nexora-api/internal/controllers"
	"nexora-api/internal/db"
	"nexora-api/internal/middlewares"
	"nexora-api/internal/repositories"
	"nexora-api/internal/routes"
	"nexora-api/internal/services"
	"nexora-api/internal/utils"
	"time"

	"github.com/gin-contrib/cors"

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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	authMiddleware := middlewares.AuthMiddleware(jwtUtil)

	routes.RegisterUserRoutes(router, authMiddleware, userController)
	routes.RegisterAuthRoutes(router, authController)

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Nexora is live ...")
	})

	router.Run(":8080")
}
