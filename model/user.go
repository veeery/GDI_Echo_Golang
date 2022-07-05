package model

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
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

type JwtCustomClaims struct{
	jwt.StandardClaims
}

func (user *User) HashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(bytes)
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func (user *User) GenerateToken() (string, error) {
	claims := &JwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))

	return t, err
}

// func ValidateToken(signedToken string) (err error) {
// 	token, err := jwt.ParseWithClaims(
// 		signedToken,
// 		&JwtCustomClaims{},
// 		func(t *jwt.Token) (interface{}, error) {
// 			return []byte(jwtKey), nil
// 		},
// 	)

// 	claims, ok := token.Claims.(*JwtCustomClaims)
// 	if !ok {
// 		err = errors.New("Cant parse claims")
// 		return
// 	}
// 	if claims.ExpiresAt < time.Now().Local().Unix() {
// 		err = errors.New("token expired")
// 		return
// 	}
// 	return
// }