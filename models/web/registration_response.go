package web

type RegistrationResponse struct {
	ID           uint   `json:"id"`
	PatientName  string `json:"patient_name"`
	DoctorName   string `json:"doctor_name"`
	ScheduleDate string `json:"schedule_date"`
	Complaint    string `json:"complaint"`
}

type CreateRegistrationResponse struct {
<<<<<<< Updated upstream
	PatientID  uint   `json:"patientId"`
	DoctorID   uint   `json:"doctorId"`
	ScheduleID uint   `json:"scheduleId"`
=======
	ID         uint   `json:"id"`
	PatientID  uint   `json:"patientID"`
	DoctorID   uint   `json:"doctorID"`
	ScheduleID uint   `json:"scheduleID"`
>>>>>>> Stashed changes
	Complaint  string `json:"complaint"`
}

type UpdateRegistrationResponse struct {
<<<<<<< Updated upstream
	PatientID  uint   `json:"patientId"`
	DoctorID   uint   `json:"doctorId"`
	ScheduleID uint   `json:"scheduleId"`
=======
	ID         uint   `json:"id"`
	PatientID  uint   `json:"patientID"`
	DoctorID   uint   `json:"doctorID"`
	ScheduleID uint   `json:"scheduleID"`
>>>>>>> Stashed changes
	Complaint  string `json:"complaint"`
}
