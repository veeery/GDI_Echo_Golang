package auth

type LogOut struct {
	Email    string `json:"email" valid:"required"`
}