package auth

// type LoginUser struct {
// 	Email    string `json:"email" valid:"required" form:"email"`
// 	Password string `json:"password" valid:"required" form:"password"`
// }

type LoginUser struct {
	Email    string `validate:"required,email" form:"email"`
	Password string `validate:"required,min=6" form:"password"`
}
