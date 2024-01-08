package model

import "time"

type Schedules struct {
	ScheduleId int
	ClassId    int
	DayOfWeek  string
	StartTime  string
	EndTime    string
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
