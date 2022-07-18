package utils

func SignedToken() string {
	a := "secret"
	return a
}

func ExpiredTokenTime() int {
	//36.000 detik = 10 jam
	i := 36000
	return i
}