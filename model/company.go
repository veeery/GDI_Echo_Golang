package model

type Company struct {
	IdCompany       uint16  `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null" json:"id_company"`
	CompanyName     string  `gorm:"type:varchar(150); not null" json:"company_name"`
	CompanyAddress  string  `gorm:"type:text;not null" json:"company_address"`
	CompanyCity     string  `gorm:"type:varchar(100);not null" json:"company_city"`
	Latitude        float64 `gorm:"type:float" json:"latitude" form:"latitude"`
	Longitude       float64 `gorm:"type:float" json:"longitude" form:"longitude"`
	CompanyPhone    string  `gorm:"type:varchar(20);not null; index" json:"company_phone"`
	CompanyLeader   string  `gorm:"type:varchar(100);not null" json:"company_leader"`
	CompanyCategory string  `gorm:"type:varchar(80);not null" json:"company_category"`
	CompanyEmail    string  `gorm:"type:varchar(80);not null; index" json:"company_email"`
}
