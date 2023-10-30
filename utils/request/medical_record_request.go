package request

import (
	"clinicai-api/models/domain"
	"clinicai-api/models/schema"
	"clinicai-api/models/web"
)

func MedicalRecordDomainToMedicalRecordSchema(request domain.MedicalRecord) *schema.MedicalRecord {
	return &schema.MedicalRecord{
		RegistrationID: request.RegistrationID,
		Symptomps:     request.Symptomps,
		Diagnoses:    request.Diagnoses,
		Solutions: request.Solutions,
	}
}

func MedicalRecordCreateRequestToMedicalRecordDomain(request web.MedicalRecordRequest) *domain.MedicalRecord {
	return &domain.MedicalRecord{
		RegistrationID: request.RegistrationID,
		Symptomps:     request.Symptomps,
		Diagnoses:    request.Diagnoses,
		Solutions: request.Solutions,
	}
}

func MedicalRecordUpdateRequestToMedicalRecordDomain(request web.MedicalRecordUpdateRequest) *domain.MedicalRecord {
	return &domain.MedicalRecord{
		RegistrationID: request.RegistrationID,
		Symptomps:     request.Symptomps,
		Diagnoses:    request.Diagnoses,
		Solutions: request.Solutions,
	}
}
