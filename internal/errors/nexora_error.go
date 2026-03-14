package nexora_error

import "net/http"

type NexoraErrorCode string

type NexoraError struct {
	ErrorCode      NexoraErrorCode `json:"errorCode"`
	ErrorMessage   string          `json:"errorMessage"`
	HttpStatusCode int             `json:"-"`
}

func (nexoraError *NexoraError) Error() string {
	return nexoraError.ErrorMessage
}

func newNexoraError(errorCode NexoraErrorCode, errorMessage string, httpStatusCode int) *NexoraError {
	return &NexoraError{
		ErrorCode:      errorCode,
		ErrorMessage:   errorMessage,
		HttpStatusCode: httpStatusCode,
	}
}

func NewInternalServerNexoraError(errorMessage string) *NexoraError {
	return &NexoraError{
		ErrorCode:      InternalServerNexoraErrorCode,
		ErrorMessage:   errorMessage,
		HttpStatusCode: http.StatusInternalServerError,
	}
}

const (
	BadRequestNexoraErrorCode         NexoraErrorCode = "NEXORA_ERROR_BAD_REQUEST"
	InternalServerNexoraErrorCode     NexoraErrorCode = "NEXORA_ERROR_INTERNAL_SERVER_ERROR"
	InvalidCredentialsNexoraErrorCode NexoraErrorCode = "NEXORA_ERROR_INVALID_CREDENTIALS"
)

var (
	BadRequestNexoraError         *NexoraError = newNexoraError(BadRequestNexoraErrorCode, "Invalid request", http.StatusBadRequest)
	InvalidCredentialsNexoraError *NexoraError = newNexoraError(InvalidCredentialsNexoraErrorCode, "Invalid credetials", http.StatusUnauthorized)
)
