package domain

type Registration struct {
	ID           uint
	PatientID    uint   `json:"patient_id"`
	DoctorID     uint   `json:"doctor_id"`
	ScheduleID   uint   `json:"schedule_id"`
	PatientName  string `json:"patient_name"`
	DoctorName   string `json:"doctor_name"`
	ScheduleDate string `json:"schedule_date"`
	Complaint    string `json:"complaint"`
}