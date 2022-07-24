package model

type Token struct {
	Token   string `json:"access_token"`
	Type    string `json:"type"`
	Expired int    `json:"expired"`
}