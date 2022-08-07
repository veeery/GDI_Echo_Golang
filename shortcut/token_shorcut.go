package shortcut

func SignedToken() string {
	a := "secret"
	return a
}

func ExpiredTokenTime() int {
	//900 detik = 15 menit
	i := 36000
	return i
}