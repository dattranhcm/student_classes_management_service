package model

import "time"

type Class struct {
	ClassId   int       `json:"class_id"`
	ClassName string    `json:"class_name"`
	TeacherId int       `json:"teacher_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
