package models

import (
	"time"
)

type Session struct {
	ID          int
	ProjectName string
	Date        string
	TimeStart   time.Time
	TimeEnd     time.Time
	ElapsedTime time.Duration
	Hours       float32
}


func (s *Session) CalculateElapsed() time.Duration {
	s.ElapsedTime = s.TimeEnd.Sub(s.TimeStart)
	return s.ElapsedTime
}
