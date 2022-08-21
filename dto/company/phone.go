package company

type Phone struct {
	// CompanyPhone string `json:"company_phone" validate:"required,numeric,min=12,max=13"`
	// IdCompanyPhone uint16 `json:"company_phone_id"`
	// IdCompany      uint16 `json:"id_company"`
	CompanyPhone string `json:"company_phone"`
}