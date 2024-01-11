package model

import "time"

type Class struct {
	ClassId   int       `json:"class_id"`
	ClassName string    `json:"class_name"`
	/*TeacherId int       `json:"teacher_id"` */
	TeacherName string `json:"teacher_name"`
	DayOfWeek string    `json:"day_of_week"`
	StartTime string    `json:"start_time"`
	EndTime   string    `json:"end_time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Students  []Users `json:"students,omitempty"`
	Teacher *Users  `json:"teacher,omitempty"`
}


type AssignStudent struct {
	StudentIds []string `json:"studentIds" validate:"required"`
}
