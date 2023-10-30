package web

type DiagnosisCreateRequest struct {
	RegistrationID uint `json:"registrationID" validate:"required"`
	OpenAIResult   string `json:"openAIResult" validate:"required"`
}

type DiagnosisUpdateRequest struct {
	RegistrationID uint `json:"registrationID" validate:"required"`
	OpenAIResult   string `json:"openAIResult" validate:"required"`
}