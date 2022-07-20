package model

import (
	"time"

	"github.com/golang-jwt/jwt"
	"gitlab.com/veeery/gdi_echo_golang.git/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	IdUser    uint32   `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null" json:"id_user"`
	FirstName string   `gorm:"type:varchar(100);not null;" json:"first_name"`
	LastName  string   `gorm:"type:varchar(100)" json:"last_name"`
	Email     string   `gorm:"type:varchar(255);not null; index" json:"email"`
	Hp        string   `gorm:"type:varchar(20);not null; index" json:"hp" `
	Password  string   `gorm:"type:varchar(100);not null" json:"-"`
	FCM 	  string   `gorm:"type:text" json:"-"`
}

type JwtCustomClaims struct{
	Email string `json:"email"`
	jwt.StandardClaims
}

func (user *User) HashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(bytes)
}

func HashPasswordUpdate(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	NewHashPassword := string(bytes)

	return NewHashPassword
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
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(utils.ExpiredTokenTime())).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(utils.SignedToken()))

	return t, err
}

func ValidateToken(signedToken string) (email string,err error) {

	tokenString := signedToken
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(utils.SignedToken()), nil
		})
	
	if token.Valid  {
		email = claims["email"].(string)
		// fmt.Println(claims["email"])
	}
	
	// for key, val := range claims {
	// 	fmt.Printf("Key: %v, value: %v\n", key, val)
	// }
	// fmt.Print(token)
	return email, err
}