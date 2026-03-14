package dto

import "nexora-api/internal/domain"

type LoginResponse struct {
	Id    string             `json:"id"`
	Email string             `json:"email"`
	Role  domain.AccountRole `json:"role"`
	Token string             `json:"token"`
}
