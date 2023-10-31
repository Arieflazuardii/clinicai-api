package web

import "time"

type ScheduleResponse struct {
<<<<<<< Updated upstream
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
=======
<<<<<<< Updated upstream
	ID       uint    	`json:"id"`
>>>>>>> Stashed changes
	DoctorID uint    	`json:"DoctorID"`
	Date     time.Time `json:"date" `
	Quota    int       `json:"quota" `
}

type ScheduleUpdateResponse struct {
	DoctorID uint   	`json:"DoctorID"`
	Date     time.Time  `json:"date" `
	Quota    int        `json:"quota"`
}
=======
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
>>>>>>> Stashed changes
