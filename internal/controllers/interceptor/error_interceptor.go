package interceptor

import (
	"fmt"
	"net/http"

	"nexora-api/internal/constants"
	nexoraError "nexora-api/internal/errors"
	nexora_error "nexora-api/internal/errors"

	"github.com/gin-gonic/gin"
)

func HandleBadRequest(context *gin.Context, err error) {
	context.AbortWithStatusJSON(
		http.StatusBadRequest,
		nexoraError.NexoraError{
			ErrorCode:    nexoraError.BadRequestNexoraErrorCode,
			ErrorMessage: err.Error(),
		},
	)
}

func HandleServiceError(context *gin.Context, nexoraError *nexoraError.NexoraError) {
	if nexoraError.HttpStatusCode == http.StatusInternalServerError {
		fmt.Printf("CRASH - %+v\n, Path: %s", nexoraError, context.Request.URL.Path)
		context.AbortWithStatusJSON(
			nexoraError.HttpStatusCode,
			nexora_error.NexoraError{
				ErrorCode:    nexoraError.ErrorCode,
				ErrorMessage: constants.INTERNAL_SERVER_ERROR_MESSAGE,
			},
		)
		return
	}

	context.AbortWithStatusJSON(nexoraError.HttpStatusCode, nexoraError)
}
