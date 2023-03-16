package domain

import "time"

type User struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Occupation   string    `json:"occupation"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	Avatar       string    `json:"avatar"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
