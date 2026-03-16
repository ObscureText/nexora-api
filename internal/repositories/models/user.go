package models

import "nexora-api/internal/domain"

type User struct {
	Id           string
	Email        string
	PasswordHash string
	Role         domain.UserRole
}
