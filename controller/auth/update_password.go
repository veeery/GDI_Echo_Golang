package auth

type UpdatePassword struct {
	Password        string `json:"password" valid:"required, length(6|100)"`
	ConfirmPassword string `json:"confirm_password" valid:"required, length(6|100)"`
}