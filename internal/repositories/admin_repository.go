package repositories

import (
	"database/sql"
	nexora_error "nexora-api/internal/errors"
	"nexora-api/internal/repositories/models"
)

type AdminRepository interface {
	GetAdminByEmail(email string) (*models.Account, *nexora_error.RepoError)
}

type adminRepository struct {
	db *sql.DB
}

func NewAdminRepository(db *sql.DB) AdminRepository {
	return &adminRepository{db: db}
}

func (adminRepository *adminRepository) GetAdminByEmail(email string) (*models.Account, *nexora_error.RepoError) {
	const query = `
		SELECT id, email, password_hash
		FROM admins
		WHERE email = $1	
	`

	var admin models.Account
	scanErr := adminRepository.db.QueryRow(query, email).Scan(
		&admin.Id,
		&admin.Email,
		&admin.PasswordHash,
	)

	if scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return nil, nexora_error.EntityNotFoundRepoError
		}

		return nil, nexora_error.NewSystemFailureRepoError("AdminRepository: GetAdminByEmail: " + scanErr.Error())
	}

	return &admin, nil
}
