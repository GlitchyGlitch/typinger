package models

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
