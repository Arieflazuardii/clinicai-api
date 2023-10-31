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
<<<<<<< Updated upstream
	PatientID  uint   `json:"patientId"`
	DoctorID   uint   `json:"doctorId"`
	ScheduleID uint   `json:"scheduleId"`
=======
=======
>>>>>>> Stashed changes
	ID         uint   `json:"id"`
	PatientID  uint   `json:"patientID"`
	DoctorID   uint   `json:"doctorID"`
	ScheduleID uint   `json:"scheduleID"`
<<<<<<< Updated upstream
>>>>>>> Stashed changes
=======
>>>>>>> Stashed changes
	Complaint  string `json:"complaint"`
}

type UpdateRegistrationResponse struct {
<<<<<<< Updated upstream
<<<<<<< Updated upstream
	PatientID  uint   `json:"patientId"`
	DoctorID   uint   `json:"doctorId"`
	ScheduleID uint   `json:"scheduleId"`
=======
=======
>>>>>>> Stashed changes
	ID         uint   `json:"id"`
	PatientID  uint   `json:"patientID"`
	DoctorID   uint   `json:"doctorID"`
	ScheduleID uint   `json:"scheduleID"`
<<<<<<< Updated upstream
>>>>>>> Stashed changes
=======
>>>>>>> Stashed changes
	Complaint  string `json:"complaint"`
}
