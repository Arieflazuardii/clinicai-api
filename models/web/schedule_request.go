package web

import "time"

type ScheduleCreateRequest struct {
	DoctorID 	uint    `json:"doctorID" form:"doctorID" validate:"required"`
	Date     	time.Time `json:"date" form:"date" validate:"required"`
	Quota 		int 	  `json:"quota" form:"quota" validate:"required"`
}

type ScheduleUpdateRequest struct {
	DoctorID 	uint    `json:"doctorID" form:"doctorID" validate:"required"`
	Date     	time.Time `json:"date" form:"date" validate:"required"`
	Quota 		int 	  `json:"quota" form:"quota" validate:"required"`
}

