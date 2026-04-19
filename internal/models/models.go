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

type Project struct {
	ID          int
	ProjectName string
}

type projectFlow struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func (s *Session) CalculateElapsed() time.Duration {
	s.ElapsedTime = s.TimeEnd.Sub(s.TimeStart)
	return s.ElapsedTime
}
