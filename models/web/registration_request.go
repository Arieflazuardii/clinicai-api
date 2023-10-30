package web

type RegistrationCreateRequest struct {
	PatientID  uint   `json:"patientID" form:"patientID" validate:"required"`
	DoctorID   uint   `json:"doctorID" form:"doctorID" validate:"required"`
	ScheduleID uint   `json:"scheduleID" form:"scheduleID" validate:"required"`
	Complaint  string `json:"complaint" form:"complaint" validate:"required"`
}

type RegistrationUpdateRequest struct {
	PatientID  uint   `json:"patientID" form:"patientID" validate:"required"`
	DoctorID   uint   `json:"doctorID" form:"doctorID" validate:"required"`
	ScheduleID uint   `json:"scheduleID" form:"scheduleID" validate:"required"`
	Complaint  string `json:"complaint" form:"complaint" validate:"required"`
}