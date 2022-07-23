package auth

type RegisterUser struct {
	FirstName       string `json:"first_name" valid:"required, alphanum" form:"first_name"`
	LastName        string `json:"last_name" form:"last_name"`
	Email           string `json:"email" valid:"required, email" form:"email"`
	Hp              string `json:"hp" valid:"required, numeric, length(12|13)" form:"hp"`
	Password        string `json:"password" valid:"required, length(6|100)" form:"password"`
	ConfirmPassword string `json:"confirm_password" valid:"required, length(6|100)" form:"confirm_password"`
}