package model

type Company struct {
	IdCompany       uint16         `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null" json:"id_company"`
	CompanyName     string         `gorm:"type:varchar(150); not null" json:"company_name" validate:"required"`
	CompanyAddress  string         `gorm:"type:text;not null" json:"company_address" validate:"required"`
	CompanyCity     string         `gorm:"type:varchar(100);not null" json:"company_city" validate:"required"`
	Latitude        float64        `gorm:"type:float;default:0.0" json:"latitude"`
	Longitude       float64        `gorm:"type:float;default:0.0" json:"longitude"`
	CompanyPhone    []CompanyPhone `gorm:"foreignKey:IdCompany" json:"detail_phone"`
	CompanyLeader   string         `gorm:"type:varchar(100);not null" json:"company_leader" validate:"required"`
	CompanyCategory string         `gorm:"type:varchar(80);not null" json:"company_category" validate:"required"`
	CompanyEmail    string         `gorm:"type:varchar(80);not null" json:"company_email" validate:"email"`
}
