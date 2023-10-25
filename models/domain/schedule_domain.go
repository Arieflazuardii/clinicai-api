package domain

import "time"

type Schedule struct {
	ID       uint
	DoctorID uint
	Date     time.Time
	Quota    int
}