package auth

// type ResetPassword struct {
// 	Password        string `json:"password" valid:"required, length(6|100)" form:"password"`
// 	ConfirmPassword string `json:"confirm_password" valid:"required, length(6|100)"`
// }

type ResetPassword struct {
	Password        string `json:"password" valid:"required, length(6|100)" form:"password"`
	ConfirmPassword string `json:"confirm_password" valid:"required, length(6|100)" form:"confirm_password"`
}