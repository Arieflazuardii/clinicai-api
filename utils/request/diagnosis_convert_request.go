package request

import (
	"clinicai-api/models/domain"
	"clinicai-api/models/schema"
	"clinicai-api/models/web"
)

func DiagnosisDomainToDiagnosisSchema(request domain.Diagnosis) *schema.Diagnosis{
	return &schema.Diagnosis{
		RegistrationID: request.RegistrationID,
		OpenAIResult: request.OpenAIResult,
	}
}

func DiagnosisCreateRequestToDiagnosisDomain(request web.DiagnosisCreateRequest) *domain.Diagnosis{
	return &domain.Diagnosis{
		RegistrationID: request.RegistrationID,
		OpenAIResult: request.OpenAIResult,
	}
}

func DiagnosisUpdateRequestToDiagnosisDomain(request web.DiagnosisUpdateRequest) *domain.Diagnosis{
	return &domain.Diagnosis{
		RegistrationID: request.RegistrationID,
		OpenAIResult: request.OpenAIResult,
	}
}
