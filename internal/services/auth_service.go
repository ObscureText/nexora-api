package services

import (
	"nexora-api/internal/controllers/dto"
	"nexora-api/internal/domain"
	nexora_error "nexora-api/internal/errors"
	"nexora-api/internal/repositories"
	"nexora-api/internal/repositories/models"
	"nexora-api/internal/utils"
)

type AuthService interface {
	Login(loginRequest *dto.LoginRequest) (*dto.LoginResponse, *nexora_error.NexoraError)
}

type authService struct {
	jwtUtil         utils.JwtUtil
	cryptoUtil      utils.CryptoUtil
	userRepository  repositories.UserRepository
	adminRepository repositories.AdminRepository
}

func NewAuthService(
	jwtUtil utils.JwtUtil,
	cryptoUtil utils.CryptoUtil,
	userRepository repositories.UserRepository,
	adminRepository repositories.AdminRepository,
) AuthService {
	return &authService{
		jwtUtil:         jwtUtil,
		cryptoUtil:      cryptoUtil,
		userRepository:  userRepository,
		adminRepository: adminRepository,
	}
}

func (authService authService) Login(loginRequest *dto.LoginRequest) (*dto.LoginResponse, *nexora_error.NexoraError) {
	account, repoErr := authService.getAccount(loginRequest)
	if repoErr != nil {
		if repoErr.ErrorCode == nexora_error.EntityNotFoundRepoErrorCode {
			return nil, nexora_error.InvalidCredentialsNexoraError
		}

		return nil, repoErr.ToNexoraInternalServerError()
	}

	isPasswordCorrect := authService.cryptoUtil.ComparePasswordAndHash(loginRequest.Password, account.PasswordHash)
	if !isPasswordCorrect {
		return nil, nexora_error.InvalidCredentialsNexoraError
	}

	token, tokenErr := authService.jwtUtil.GenerateToken(
		account.Id,
		loginRequest.UserRole,
	)

	if tokenErr != nil {
		return nil, nexora_error.NewInternalServerNexoraError("AuthService: Login: GenerateToken error: " + tokenErr.Error())
	}

	return &dto.LoginResponse{
		Id:    account.Id,
		Email: account.Email,
		Role:  loginRequest.UserRole,
		Token: token,
	}, nil
}

func (authService authService) getAccount(loginRequest *dto.LoginRequest) (*models.Account, *nexora_error.RepoError) {
	if loginRequest.UserRole == domain.Admin {
		return authService.adminRepository.GetAdminByEmail(loginRequest.Email)
	}

	if loginRequest.UserRole == domain.Individual || loginRequest.UserRole == domain.CompanyOwner {
		return authService.userRepository.GetUserByEmailAndRole(
			loginRequest.Email,
			string(loginRequest.UserRole),
		)
	}

	return nil, nexora_error.EntityNotFoundRepoError
}
