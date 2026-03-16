package dto

import "nexora-api/internal/domain"

type LoginResponse struct {
	Token string          `json:"token"`
	Role  domain.UserRole `json:"role"`
}
