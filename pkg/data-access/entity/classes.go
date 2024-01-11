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
	Teacher       *User     `bun:"rel:belongs-to,join:teacher_id=user_id"`
	DayOfWeek     string    `bun:"day_of_week"`
	StartTime     string    `bun:"start_time"`
	EndTime       string    `bun:"end_time"`
	CreatedAt     time.Time `bun:"created_at,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at,default:current_timestamp"`
}
