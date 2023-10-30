package web

type RegistrationResponse struct {
	ID           uint   `json:"id"`
	PatientName  string `json:"patient_name"`
	DoctorName   string `json:"doctor_name"`
	ScheduleDate string `json:"schedule_date"`
	Complaint    string `json:"complaint"`
}

type CreateRegistrationResponse struct {
	PatientID  uint   `json:"patientId"`
	DoctorID   uint   `json:"doctorId"`
	ScheduleID uint   `json:"scheduleId"`
	Complaint  string `json:"complaint"`
}

type UpdateRegistrationResponse struct {
	PatientID  uint   `json:"patientId"`
	DoctorID   uint   `json:"doctorId"`
	ScheduleID uint   `json:"scheduleId"`
	Complaint  string `json:"complaint"`
}
