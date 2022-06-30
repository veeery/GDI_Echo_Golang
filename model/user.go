package model

import (
	"os"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	IdUser   int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null" json:"id_user"`
	Name     string `gorm:"type:varchar(225)" json:"name"`
	Email    string `gorm:"type:varchar(255)" json:"email"`
	// Hp       string `gorm:"type:varchar(20)" json:"hp"`
	Password string `gorm:"->;<-;not null" json:"-"`
	// Token string `gorm:"-" json:"token,omitempty"`
}

var (
	jwtKey = os.Getenv("JWT_KEY")
)

func (user *User) HashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(bytes)
}

func (user *User) GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"IdUser": user.IdUser,
	})

	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}
