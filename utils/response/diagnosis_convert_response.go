package response

import (
	"clinicai-api/models/domain"
	"clinicai-api/models/schema"
	"clinicai-api/models/web"
)

func DiagnosisSchemaToDiagnosisDomain(dianogsis *schema.Diagnosis) *domain.Diagnosis {
	return &domain.Diagnosis{
		ID:				dianogsis.ID,
		RegistrationID: dianogsis.RegistrationID,
		OpenAIResult:   dianogsis.OpenAIResult,
	}
}

func DiagnosisDomainToDiagnosisResponse(dianogsis *domain.Diagnosis, registration *domain.Registration) web.DiagnosisResponse {
	return web.DiagnosisResponse{
		RegistrationID: dianogsis.RegistrationID,
		Registration: registration,
		OpenAIResult:   dianogsis.OpenAIResult,
	}
}


func CreateDiagnosisDomainToDiagnosisResponse(dianogsis *domain.Diagnosis) web.DiagnosisResponse {
	return web.DiagnosisResponse{
		RegistrationID: dianogsis.RegistrationID,
		OpenAIResult:   dianogsis.OpenAIResult,
	}
}

func ConvertDiagnosisResponse(diagnostic []domain.Diagnosis) []web.DiagnosisResponse {
	var results []web.DiagnosisResponse
	for _, dianogsis := range diagnostic {
		dianogsisResponse := web.DiagnosisResponse{
			RegistrationID: dianogsis.RegistrationID,
			OpenAIResult: dianogsis.OpenAIResult,
		}
		results = append(results, dianogsisResponse)
	}
	return results
}
