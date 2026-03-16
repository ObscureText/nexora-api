package middlewares

import (
	"errors"
	"net/http"
	"nexora-api/internal/constants"
	nexora_error "nexora-api/internal/errors"
	"nexora-api/internal/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(jwtUtil utils.JwtUtil) gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, nexora_error.UnauthorizedNexoraError)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, nexora_error.UnauthorizedNexoraError)
			return
		}

		token := parts[1]

		userId, userRole, err := jwtUtil.ParseToken(token)
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				context.AbortWithStatusJSON(http.StatusUnauthorized, nexora_error.TokenExpiredNexoraError)
				return
			}

			context.AbortWithStatusJSON(http.StatusUnauthorized, nexora_error.UnauthorizedNexoraError)
			return
		}

		context.Set(constants.KEY_USER_ID, userId)
		context.Set(constants.KEY_USER_ROLE, userRole)

		context.Next()
	}
}
