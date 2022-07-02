package auth

type RegisterUser struct {
	FirstName string `json:"first_name" valid:"required, alphanum"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" valid:"required, email"`
	Hp        string `json:"hp" valid:"required, numeric, length(12|13)"`
	Password  string `json:"password" valid:"required, length(6|100)"`
}