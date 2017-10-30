package model

type User struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	PhoneModel string `json:"phoneModel"`
}

type Token struct {
	Token string `json:"token"`
}
