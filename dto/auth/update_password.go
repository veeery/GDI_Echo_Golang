package auth

type ResetPassword struct {
	Password        string `json:"password" validate:"required,min=6" form:"password"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=6" form:"confirm_password"`
}