package services

import (
	"nexora-api/internal/controllers/dto"
	nexora_error "nexora-api/internal/errors"
	"nexora-api/internal/repositories"
)

type UserService interface {
	GetUser(userId string) (*dto.UserDetailsResponse, *nexora_error.NexoraError)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}
func (service *userService) GetUser(userId string) (*dto.UserDetailsResponse, *nexora_error.NexoraError) {
	user, serviceProvider, repoErr := service.userRepository.GetUserDetailsById(userId)
	if repoErr != nil {
		if repoErr.ErrorCode == nexora_error.UserNotFoundRepoErrorCode {
			return nil, nexora_error.UnauthorizedNexoraError
		}

		return nil, repoErr.ToNexoraInternalServerError()
	}

	response := &dto.UserDetailsResponse{
		User: dto.User{
			Id:    user.Id,
			Email: user.Email,
			Role:  user.Role,
		},
		ServiceProvider: dto.ServiceProvider{
			Id:           serviceProvider.Id,
			Name:         serviceProvider.Name,
			Email:        serviceProvider.Email,
			MobileNumber: serviceProvider.MobileNumber,
			Address:      serviceProvider.Address,
		},
	}

	return response, nil
}
