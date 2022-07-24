package auth

type LoginUser struct {
	Email    string `validate:"required,email" form:"email"`
	Password string `validate:"required,min=6" form:"password"`
}
