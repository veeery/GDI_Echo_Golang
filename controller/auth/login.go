package auth

type LoginUser struct {
	Email    string `json:"email" valid:"required"`
	Password string `json:"password" valid:"required"`
}