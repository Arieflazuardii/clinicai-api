package web

import "clinicai-api/models/domain"

type DiagnosisResponse struct {
	RegistrationID uint                `json:"registrationID"`
	Registration   *domain.Registration `json:"registration"`
	OpenAIResult   string              `json:"openAIResult"`
}

type DiagnosisCreateResponse struct {
	RegistrationID uint   `json:"registrationID"`
	OpenAIResult   string `json:"openAIResult"`
}

type DiagnosisUpdateResponse struct {
	RegistrationID uint   `json:"registrationID"`
	OpenAIResult   string `json:"openAIResult"`
}
