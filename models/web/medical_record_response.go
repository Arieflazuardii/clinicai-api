package web

import "clinicai-api/models/domain"

type MedicalRecordResponse struct {
	ID             uint                 `json:"id"`
	RegistrationID uint                 `json:"registrationID"`
	Registration   *domain.Registration `json:"registration"`
	Symptomps      string               `json:"symptomps"`
	Diagnoses      string               `json:"diagnoses"`
	PatientName string `json:"patient_name"`
	DoctorName string `json:"doctor_name"`
	Solutions      string               `json:"solutions"`
}

type ConvertMedicalRecordResponse struct {
	ID             uint                 `json:"id"`
	RegistrationID uint                 `json:"registrationID"`
	Symptomps      string               `json:"symptomps"`
	Diagnoses      string               `json:"diagnoses"`
	Solutions      string               `json:"Solutions"`
	PatientName    string               `json:"patient_name"`
	DoctorName     string               `json:"doctor_name"`
}

type MedicalRecordCreateResponse struct {
	ID             uint   `json:"id"`
	RegistrationID uint   `json:"RegistrationID"`
	Symptomps      string `json:"symptomps"`
	Diagnoses      string `json:"diagnoses"`
	Solutions      string `json:"solutions"`
}

type MedicalRecordUpdateResponse struct {
	ID             uint   `json:"id"`
	RegistrationID uint   `json:"RegistrationID"`
	Symptomps      string `json:"symptomps"`
	Diagnoses      string `json:"diagnoses"`
	Solutions      string `json:"solutions"`
}