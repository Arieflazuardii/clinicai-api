package domain

type MedicalRecord struct {
	ID             uint
	RegistrationID uint
	Registration   Registration
	Symptomps      string
	Diagnoses      string
	Solutions      string
	PatientName    string
	DoctorName     string
}