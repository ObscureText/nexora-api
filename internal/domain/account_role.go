package domain

type UserRole string

const (
	Individual           UserRole = "INDIVIDUAL"
	CompanyOwner         UserRole = "ESTABLISHMENT_OWNER"
	ServiceProviderAdmin UserRole = "SERVICE_PROVIDER_ADMIN"
)
