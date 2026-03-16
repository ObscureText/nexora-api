package dto

import "nexora-api/internal/domain"

type UserDetailsResponse struct {
	User            User            `json:"user"`
	ServiceProvider ServiceProvider `json:"service_provider"`
}

type User struct {
	Id    string          `json:"id"`
	Email string          `json:"email"`
	Role  domain.UserRole `json:"role"`
}

type ServiceProvider struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	MobileNumber string `json:"mobile_number"`
	Address      string `json:"address"`
}
