package web

import "time"

type ScheduleResponse struct {
	ID uint `json:"id"`
	DoctorID uint    	`json:"doctorID"`
	Date     time.Time `json:"date"`
	Quota    int       `json:"quota"`
	DoctorName string `json:"doctor_name"`
}

type ConvertScheduleResponse struct {
	ID uint `json:"id"`
	DoctorID uint    	`json:"doctorID"`
	Date     time.Time `json:"date"`
	Quota    int       `json:"quota"`
	DoctorName string `json:"doctor_name"`
}

type ScheduleUpdateResponse struct {
	ID uint `json:"id"`
	DoctorID uint    	`json:"DoctorID"`
	Date     time.Time `json:"date"`
	Quota    int       `json:"quota"`
	DoctorName string `json:"doctor_name"`
}
