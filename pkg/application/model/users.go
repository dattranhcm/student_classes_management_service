package model

import "time"

type Users struct {
	UserId    int       `json:"user_id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	UserType  string    `json:"user_type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
