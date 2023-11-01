package model

type UserCredentials struct {
	ID       string
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLogout struct {
	Email string
}
