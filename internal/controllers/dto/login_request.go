package dto

import "nexora-api/internal/domain"

type LoginRequest struct {
	Email    string             `json:"email" binding:"required,email"`
	Password string             `json:"password" binding:"required,min=6"`
	UserRole domain.AccountRole `json:"user_role" binding:"required"`
}
