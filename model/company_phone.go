package model

type CompanyPhone struct {
	IdCompanyPhone uint16 `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null" json:"company_phone_id"`
	IdCompany      uint16 `json:"id_company"`
	CompanyPhone   string `gorm:"type:varchar(20);not null" json:"company_phone" validate:"required,min=10,max=12"`
}
