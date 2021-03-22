package models

import "time"

type User struct {
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	Email        string     `json:"email"`
	PasswordHash string     `json:"password_hash"`
	DeletedAt    *time.Time `json:"-" pg:",soft_delete"` // TODO: test soft_delete
}

type NewUser struct {
	Name     string `json:"name" validate:"min=2,max=64"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=8,max=256"`
}

type UpdateUser struct {
	Name     string `json:"name" validate:"omitempty,min=2,max=64"`
	Email    string `json:"email" validate:"omitempty,email"`
	Password string `json:"password" validate:"omitempty,min=8,max=256"`
}

type UserFilter struct {
	Name  string `json:"name" validate:"omitempty,max=64"`
	Email string `json:"email" validate:"omitempty,max=320"`
}
