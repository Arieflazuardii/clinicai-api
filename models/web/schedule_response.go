package web

import "time"

type ScheduleResponse struct {
	ID       uint    	`json:"id"`
	DoctorID uint    	`json:"DoctorID"`
	Date     time.Time `json:"date" `
	Quota    int       `json:"quota" `
}

type ScheduleUpdateResponse struct {
	ID       uint    	`json:"id"`
	DoctorID uint   	`json:"DoctorID"`
	Date     time.Time  `json:"date" `
	Quota    int        `json:"quota"`
}
