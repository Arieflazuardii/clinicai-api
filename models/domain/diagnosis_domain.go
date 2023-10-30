package domain

type Diagnosis struct {
	ID             uint
	RegistrationID uint
	Registration   Registration
	OpenAIResult   string
	PatientName    string
	DoctorName     string
	ScheduleDate   string
	Complaint      string
}