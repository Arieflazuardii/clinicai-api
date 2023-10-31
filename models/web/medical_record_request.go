package web

type MedicalRecordRequest struct {
	RegistrationID uint    	`json:"RegistrationID"`
	Symptomps string `json:"symptomps"`
	Diagnoses string `json:"diagnoses"`
	Solutions string `json:"solutions"`
}

type MedicalRecordUpdateRequest struct {
	RegistrationID uint    	`json:"RegistrationID"`
	Symptomps string `json:"symptomps"`
	Diagnoses string `json:"diagnoses"`
	Solutions string `json:"solutions"`
}


