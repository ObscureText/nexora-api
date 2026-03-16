package repositories

import (
	"database/sql"
	nexora_error "nexora-api/internal/errors"
	"nexora-api/internal/repositories/models"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, *nexora_error.RepoError)
	GetUserDetailsById(userId string) (*models.User, *models.ServiceProvider, *nexora_error.RepoError)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (userRepository *userRepository) GetUserByEmail(email string) (*models.User, *nexora_error.RepoError) {
	const query = `
		SELECT users.id, users.email, users.password_hash, user_roles.name
		FROM users
		JOIN user_roles ON user_roles.id = users.role_id
		WHERE users.email = $1;
	`

	var user models.User

	scanErr := userRepository.db.QueryRow(query, email).Scan(
		&user.Id,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
	)

	if scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return nil, nexora_error.UserNotFoundRepoError
		}

		return nil, nexora_error.NewSystemFailureRepoError("UserRepository: GetUserByEmail: " + scanErr.Error())
	}

	return &user, nil
}

func (repo *userRepository) GetUserDetailsById(userId string) (*models.User, *models.ServiceProvider, *nexora_error.RepoError) {
	const query = `
		SELECT
			users.id,
			users.email,
			user_roles.name,

			service_providers.id,
			service_providers.name,
			service_providers.email,
			service_providers.mobile_number,
			service_providers.address
		FROM users
		JOIN user_roles ON user_roles.id = users.role_id
		JOIN service_providers ON service_providers.id = users.service_provider_id
		WHERE users.id = $1;
	`

	var user models.User
	var serviceProvider models.ServiceProvider

	scanErr := repo.db.QueryRow(query, userId).Scan(
		&user.Id,
		&user.Email,
		&user.Role,

		&serviceProvider.Id,
		&serviceProvider.Name,
		&serviceProvider.Email,
		&serviceProvider.MobileNumber,
		&serviceProvider.Address,
	)

	if scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return nil, nil, nexora_error.UserNotFoundRepoError
		}

		return nil, nil, nexora_error.NewSystemFailureRepoError("UserRepository: GetUserDetailsById: " + scanErr.Error())
	}

	return &user, &serviceProvider, nil
}
