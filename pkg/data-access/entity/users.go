package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`
	UserId        int       `bun:"user_id,pk,autoincrement"`
	Username      string    `bun:"username"`
	Password      string    `bun:"password"`
	UserType      string    `bun:"user_type"`
	CreatedAt     time.Time `bun:"created_at,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at,default:current_timestamp"`
}
