package web

type RegistrationResponse struct {
	ID           uint   `json:"id"`
	PatientName  string `json:"patient_name"`
	DoctorName   string `json:"doctor_name"`
	ScheduleDate string `json:"schedule_date"`
	Complaint    string `json:"complaint"`
}

type CreateRegistrationResponse struct {
	ID         uint   `json:"id"`
	PatientID  uint   `json:"patientID"`
	DoctorID   uint   `json:"doctorID"`
	ScheduleID uint   `json:"scheduleID"`
	Complaint  string `json:"complaint"`
}

type UpdateRegistrationResponse struct {
	ID         uint   `json:"id"`
	PatientID  uint   `json:"patientID"`
	DoctorID   uint   `json:"doctorID"`
	ScheduleID uint   `json:"scheduleID"`
	Complaint  string `json:"complaint"`
}
