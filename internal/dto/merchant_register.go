// pkg/handler/dto/merchant_register.go
package dto

type AddressDTO struct {
	Street     string `json:"street"`
	District   string `json:"district"`
	Province   string `json:"province"`
	PostalCode string `json:"postalCode"`
	Country    string `json:"country"` // optional; defaults to Thailand
}

type MerchantRegisterDTO struct {
	// Tenant fields
	LegalName  string     `json:"legalName"`
	BranchCode string     `json:"branchCode"`
	TaxID      string     `json:"taxId"`
	Address    AddressDTO `json:"address"`
	LogoURL    string     `json:"logoUrl"`

	// Owner-user fields
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
}
