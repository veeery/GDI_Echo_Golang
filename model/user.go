package model

// type User struct {
// 	// gorm.Model
// 	Id       int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null" json:"Id"`
// 	Name     string `gorm:"type:varchar(225)" json:"name"`
// 	Email    string `gorm:"type:varchar(255)" json:"email"`
// 	Hp       string `gorm:"type:varchar(20)" json:"hp"`
// 	Password string `gorm:"->;<-;not null" json:"-"`
// 	Token    string `gorm:"-" json:"token,omitempty"`
// }

type User struct {
	ID       uint64  `gorm:"primary_key:auto_increment" json:"id"`
	Name     string  `gorm:"type:varchar(255)" json:"name"`
	Email    string  `gorm:"type:varchar(255)" json:"email"`
	Password string  `gorm:"->;<-;not null" json:"-"`
	Token    string  `gorm:"-" json:"token,omitempty"`
}