package auth

type RefreshUser struct {
	IdUser uint16 `json:"id_user"`
	Email  string `json:"email"`
}
