package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type Class struct {
	bun.BaseModel `bun:"table:classes"`
	ClassId       int       `bun:"class_id,pk,autoincrement"`
	ClassName     string    `bun:"class_name"`
	TeacherId     int       `bun:"teacher_id"`
	User          User      `bun:"rel:belongs-to,join:teacher_id=user_id"`
	CreatedAt     time.Time `bun:"created_at,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at,default:current_timestamp"`
}
