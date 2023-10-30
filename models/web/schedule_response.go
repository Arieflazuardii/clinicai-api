package web

import "time"

type ScheduleResponse struct {
	DoctorID uint    	`json:"doctorID"`
	Date     time.Time `json:"date"`
	Quota    int       `json:"quota"`
	DoctorName string `json:"doctor_name"`
}

type ConvertScheduleResponse struct {
	DoctorID uint    	`json:"doctorID"`
	Date     time.Time `json:"date"`
	Quota    int       `json:"quota"`
	DoctorName string `json:"doctor_name"`
}

type ScheduleUpdateResponse struct {
	DoctorID uint   	`json:"DoctorID"`
	Date     time.Time  `json:"date" `
	Quota    int        `json:"quota"`
}


