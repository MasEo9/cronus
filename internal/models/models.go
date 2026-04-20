package models

// domain models, i.e. anything related to database structure

import (
	"time"
)

type Session struct {
	ID          int
	ProjectID   int
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

func (s *Session) CalculateElapsed() time.Duration {
	s.ElapsedTime = s.TimeEnd.Sub(s.TimeStart)
	return s.ElapsedTime
}

