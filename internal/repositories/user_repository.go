package repositories

import (
	"database/sql"
	nexora_error "nexora-api/internal/errors"
	"nexora-api/internal/repositories/models"
)

type UserRepository interface {
	GetUserByEmailAndRole(email string, role string) (*models.Account, *nexora_error.RepoError)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (userRepository *userRepository) GetUserByEmailAndRole(email string, role string) (*models.Account, *nexora_error.RepoError) {
	const query = `
		SELECT users.id, users.email, users.password_hash
		FROM users
		JOIN user_roles ON user_roles.id = users.role_id
		WHERE users.email = $1 AND user_roles.name = $2;	
	`

	var user models.Account
	scanErr := userRepository.db.QueryRow(query, email, role).Scan(
		&user.Id,
		&user.Email,
		&user.PasswordHash,
	)

	if scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return nil, nexora_error.EntityNotFoundRepoError
		}

		return nil, nexora_error.NewSystemFailureRepoError("UserRepository: GetUserByEmailAndRole: " + scanErr.Error())
	}

	return &user, nil
}
