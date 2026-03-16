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
	UserNotFoundRepoErrorCode  RepoErrorCode = "REPO_ERROR_USER_NOT_FOUND"
	SystemFailureRepoErrorCode RepoErrorCode = "REPO_ERROR_SYSTEM_FAILURE"
)

var (
	UserNotFoundRepoError *RepoError = newRepoError(UserNotFoundRepoErrorCode, "User Not Found")
)
