package model

import (
	"os"

	// jwt "github.com/dgrijalva/jwt-go"

	jwt "github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var (
	jwtKey = os.Getenv("JWT_KEY")
)

type User struct {
	IdUser    uint32   `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null" json:"id_user"`
	FirstName string   `gorm:"type:varchar(100);not null;" json:"first_name"`
	LastName  string   `gorm:"type:varchar(100)" json:"last_name"`
	Email     string   `gorm:"type:varchar(255);not null; index" json:"email"`
	Hp        string   `gorm:"type:varchar(20);not null; index" json:"hp" `
	Password  string   `gorm:"type:varchar(100);not null" json:"-"`
	// Token    string   `gorm:"-" json:"token,omitempty"`
}


func (user *User) HashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(bytes)
}

func (user *User) GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Email": user.Email,
	})

	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}
