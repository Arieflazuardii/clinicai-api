package web

import "clinicai-api/models/domain"

type DiagnosisResponse struct {
	ID uint `json:"id"`
	RegistrationID uint                `json:"registrationID"`
	Registration   *domain.Registration `json:"registration"`
	OpenAIResult   string              `json:"openAIResult"`
}

type DiagnosisCreateResponse struct {
	ID uint `json:"id"`
	RegistrationID uint   `json:"registrationID"`
	OpenAIResult   string `json:"openAIResult"`
}

type DiagnosisUpdateResponse struct {
	ID uint `json:"id"`
	RegistrationID uint   `json:"registrationID"`
	OpenAIResult   string `json:"openAIResult"`
}
