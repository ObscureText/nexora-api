package services

import (
	"nexora-api/internal/controllers/dto"
	nexora_error "nexora-api/internal/errors"
	"nexora-api/internal/repositories"
	"nexora-api/internal/utils"
)

type AuthService interface {
	Login(loginRequest *dto.LoginRequest) (*dto.LoginResponse, *nexora_error.NexoraError)
}

type authService struct {
	jwtUtil        utils.JwtUtil
	cryptoUtil     utils.CryptoUtil
	userRepository repositories.UserRepository
}

func NewAuthService(
	jwtUtil utils.JwtUtil,
	cryptoUtil utils.CryptoUtil,
	userRepository repositories.UserRepository,
) AuthService {
	return &authService{
		jwtUtil:        jwtUtil,
		cryptoUtil:     cryptoUtil,
		userRepository: userRepository,
	}
}

func (authService authService) Login(loginRequest *dto.LoginRequest) (*dto.LoginResponse, *nexora_error.NexoraError) {
	user, repoErr := authService.userRepository.GetUserByEmail(loginRequest.Email)
	if repoErr != nil {
		if repoErr.ErrorCode == nexora_error.UserNotFoundRepoErrorCode {
			return nil, nexora_error.InvalidCredentialsNexoraError
		}

		return nil, repoErr.ToNexoraInternalServerError()
	}

	isPasswordCorrect := authService.cryptoUtil.ComparePasswordAndHash(loginRequest.Password, user.PasswordHash)
	if !isPasswordCorrect {
		return nil, nexora_error.InvalidCredentialsNexoraError
	}

	token, tokenErr := authService.jwtUtil.GenerateToken(
		user.Id,
		user.Role,
	)

	if tokenErr != nil {
		return nil, nexora_error.NewInternalServerNexoraError("AuthService: Login: GenerateToken error: " + tokenErr.Error())
	}

	return &dto.LoginResponse{Token: token}, nil
}
