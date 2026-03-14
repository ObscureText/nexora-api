package nexora_error

import "net/http"

type RepoErrorCode string

type RepoError struct {
	ErrorCode    RepoErrorCode `json:"errorCode"`
	ErrorMessage string        `json:"errorMessage"`
}

func (nexoraError *RepoError) Error() string {
	return nexoraError.ErrorMessage
}

func (nexoraError *RepoError) ToNexoraInternalServerError() *NexoraError {
	return &NexoraError{
		ErrorCode:      InternalServerNexoraErrorCode,
		ErrorMessage:   nexoraError.ErrorMessage,
		HttpStatusCode: http.StatusInternalServerError,
	}
}

func newRepoError(errorCode RepoErrorCode, errorMessage string) *RepoError {
	return &RepoError{
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
	}
}

func NewSystemFailureRepoError(errorMessage string) *RepoError {
	return &RepoError{
		ErrorCode:    SystemFailureRepoErrorCode,
		ErrorMessage: errorMessage,
	}
}

const (
	SystemFailureRepoErrorCode  RepoErrorCode = "REPO_ERROR_SYSTEM_FAILURE"
	EntityNotFoundRepoErrorCode RepoErrorCode = "REPO_ERROR_ENTITY_NOT_FOUND"
)

var (
	EntityNotFoundRepoError *RepoError = newRepoError(EntityNotFoundRepoErrorCode, "Entity Not Found")
)
