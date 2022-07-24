package company

type RegisterCompany struct {
	CompanyName     string  `json:"company_name" validate:"required" form:"company_name"`
	CompanyAddress  string  `json:"company_address" validate:"required" form:"company_address"`
	CompanyCity     string  `json:"company_city" validate:"required" form:"company_city"`
	Latitude        float64 `json:"latitude" form:"latitude"`
	Longitude       float64 `json:"longitude" form:"longitude"`
	CompanyPhone    string  `json:"company_phone" validate:"required,numeric,min=12,max=13" form:"company_phone"`
	CompanyLeader   string  `json:"company_leader" validate:"required" form:"company_leader"`
	CompanyCategory string  `json:"company_category" validate:"required" form:"company_category"`
	CompanyEmail    string  `json:"company_email" validate:"required,email" form:"company_email"`
}