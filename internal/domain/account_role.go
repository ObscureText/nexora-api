package domain

type AccountRole string

const (
	Admin        AccountRole = "ADMIN"
	Individual   AccountRole = "INDIVIDUAL"
	CompanyOwner AccountRole = "COMPANY_OWNER"
)

func (role AccountRole) IsValid() bool {
	return role == CompanyOwner || role == Individual || role == Admin
}
