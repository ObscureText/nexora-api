package middlewares

import (
	"net/http"
	"nexora-api/internal/constants"
	"nexora-api/internal/domain"
	nexora_error "nexora-api/internal/errors"
	"slices"

	"github.com/gin-gonic/gin"
)

func RequireRoles(allowedRoles ...domain.UserRole) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		roleValue, exists := ctx.Get(constants.KEY_USER_ROLE)
		if !exists {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, nexora_error.UnauthorizedNexoraError)
			return
		}

		userRole, ok := roleValue.(domain.UserRole)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, nexora_error.UnauthorizedNexoraError)
			return
		}

		if slices.Contains(allowedRoles, userRole) {
			ctx.Next()
			return
		}

		ctx.AbortWithStatusJSON(http.StatusForbidden, nexora_error.UnauthorizedNexoraError)
	}
}
