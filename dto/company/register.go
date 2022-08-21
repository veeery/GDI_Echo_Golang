package company

type RegisterCompany struct {
	CompanyName    string `json:"company_name" validate:"required"`
	CompanyAddress string `json:"company_address" validate:"required"`
	// CompanyCity    string  `json:"company_city" validate:"required"`
	// Latitude       float64 `json:"latitude"`
	// Longitude      float64 `json:"longitude"`
	CompanyPhone []Phone `json:"detail_phone" validate:"required"`

	// CompanyLeader   string `json:"company_leader" validate:"required"`
	// CompanyCategory string `json:"company_category" validate:"required"`
	// CompanyEmail    string `json:"company_email" validate:"required,email"`
}