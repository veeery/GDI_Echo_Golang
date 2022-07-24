package auth

type RegisterUser struct {
	FirstName       string `json:"first_name" validate:"required" form:"first_name"`
	LastName        string `json:"last_name" form:"last_name"`
	Email           string `json:"email" validate:"required,email" form:"email"`
	Hp              string `json:"hp" validate:"required,numeric,min=12,max=13" form:"hp"`
	Password        string `json:"password" validate:"required,min=6" form:"password"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=6" form:"confirm_password"`
}