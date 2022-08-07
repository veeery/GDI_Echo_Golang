package company

type Phone struct {
	CompanyPhone string `json:"company_phone" validate:"required,numeric,min=12,max=13" form:"company_phone"`
	// Provider     string `json:"provider" form:"provider"`
}