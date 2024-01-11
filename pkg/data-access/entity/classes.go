package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type (
	Class struct {
		bun.BaseModel `bun:"table:classes"`
		ClassId       int       `bun:"class_id,pk,autoincrement"`
		ClassName     string    `bun:"class_name"`
		TeacherId     int       `bun:"teacher_id"`
		DayOfWeek     string    `bun:"day_of_week"`
		StartTime     string    `bun:"start_time"`
		EndTime       string    `bun:"end_time"`
		CreatedAt     time.Time `bun:"created_at,default:current_timestamp"`
		UpdatedAt     time.Time `bun:"updated_at,default:current_timestamp"`
		Teacher       *User     `bun:"rel:belongs-to,join:teacher_id=user_id"`
		Students      []User 	`bun:"m2m:students_classes,join:Class=Student"`
	}

	StudentClass struct {
		bun.BaseModel `bun:"table:students_classes"`
		ClassID       string   `bun:"class_id,pk"`
		Class         *Class   `bun:"rel:belongs-to,join:class_id=class_id"`
		StudentID     string   `bun:"student_id,pk"`
		Student       *User `bun:"rel:belongs-to,join:student_id=user_id"`
	}
)
