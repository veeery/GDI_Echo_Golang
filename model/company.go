package model

type Company struct {
	IdCompany      uint16         `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null" json:"id_company"`
	CompanyName    string         `gorm:"type:varchar(150); not null" json:"company_name"`
	CompanyAddress string         `gorm:"type:text;not null" json:"company_address"`
	CompanyCity    string         `gorm:"type:varchar(100);not null" json:"company_city"`
	Latitude       float64        `gorm:"type:float" json:"latitude" json:"latitude"`
	Longitude      float64        `gorm:"type:float" json:"longitude" json:"longitude"`
	CompanyPhone   []CompanyPhone `gorm:"foreignKey:IdCompany" json:"detail_phone" form:"detail_phone"`
	// CompanyPhone    []CompanyPhone `gorm:"embedded_prefix:" json:"list_company_phone"`
	CompanyLeader   string `gorm:"type:varchar(100);not null" json:"company_leader"`
	CompanyCategory string `gorm:"type:varchar(80);not null" json:"company_category"`
	CompanyEmail    string `gorm:"type:varchar(80);not null" json:"company_email"`
}
