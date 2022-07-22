package auth

type LoginUser struct {
	Email    string `json:"email" valid:"required" form:"email"`
	Password string `json:"password" valid:"required" form:"password"`
}
