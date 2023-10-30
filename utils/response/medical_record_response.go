package response

import (
	"clinicai-api/models/domain"
	"clinicai-api/models/schema"
	"clinicai-api/models/web"
)

func MedicalRecordSchemaToMedicalRecordDomain(medicalRecord *schema.MedicalRecord) *domain.MedicalRecord {
	return &domain.MedicalRecord{
		ID: medicalRecord.ID,
		RegistrationID: medicalRecord.RegistrationID,
		Symptomps:     medicalRecord.Symptomps,
		Diagnoses:    medicalRecord.Diagnoses,
		Solutions: medicalRecord.Solutions,
	}
}

func MedicalRecordDomainToMedicalRecordResponse(medicalRecord *domain.MedicalRecord, registration *domain.Registration) web.MedicalRecordResponse {
	return web.MedicalRecordResponse{
		ID: 	  medicalRecord.ID,
		RegistrationID: medicalRecord.RegistrationID,
		Registration: registration,
		Symptomps:     medicalRecord.Symptomps,
		Diagnoses:    medicalRecord.Diagnoses,
		Solutions: medicalRecord.Solutions,
	}
}

func CreateMedicalRecordDomainToMedicalRecordResponse(medicalRecord *domain.MedicalRecord) web.MedicalRecordResponse {
	return web.MedicalRecordResponse{
		ID: 	  medicalRecord.ID,
		RegistrationID: medicalRecord.RegistrationID,
		Symptomps:     medicalRecord.Symptomps,
		Diagnoses:    medicalRecord.Diagnoses,
		Solutions: medicalRecord.Solutions,
	}
}

func UpdateMedicalRecordDomainToMedicalRecordResponse(medicalRecord *domain.MedicalRecord) web.MedicalRecordUpdateResponse {
	return web.MedicalRecordUpdateResponse{
		ID: 	  medicalRecord.ID,
		RegistrationID: medicalRecord.RegistrationID,
		Symptomps:     medicalRecord.Symptomps,
		Diagnoses:    medicalRecord.Diagnoses,
		Solutions: medicalRecord.Solutions,
	}
}

func ConvertMedicalRecordResponse(medicalRecords []domain.MedicalRecord) []web.ConvertMedicalRecordResponse {
	var results []web.ConvertMedicalRecordResponse
	for _, medicalRecord := range medicalRecords {
		medicalRecordResponse := web.ConvertMedicalRecordResponse {
			ID: 	  medicalRecord.ID,
			RegistrationID: medicalRecord.RegistrationID,
			PatientName: medicalRecord.PatientName,
			DoctorName: medicalRecord.DoctorName,
			Symptomps:     medicalRecord.Symptomps,
			Diagnoses:    medicalRecord.Diagnoses,
			Solutions: medicalRecord.Solutions,
		}
		results = append(results, medicalRecordResponse)
	}
	return results
}
