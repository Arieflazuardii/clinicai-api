package response

import (
	"clinicai-api/models/domain"
	"clinicai-api/models/schema"
	"clinicai-api/models/web"
)

func PatientDomainToPatientLoginResponse(patient *domain.Patient) web.PatientLoginResponse {
	return web.PatientLoginResponse{
		Name:   patient.Name,
		Email:  patient.Email,
	}
}

func PatientSchemaToPatientDomain(patient *schema.Patient) *domain.Patient {
	return &domain.Patient{
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
		ID:           patient.ID,
=======
		ID: patient.ID,
>>>>>>> Stashed changes
>>>>>>> Stashed changes
		Name: 	      patient.Name,
		Email:        patient.Email,
		Nik:          patient.Nik,
		Birthday:     patient.Birthday,
		Age:          patient.Age,
		Address:      patient.Address,
		Gender:       patient.Gender,
		Phone_number: patient.Phone_number,
	}
}

func PatientDomainToPatientResponse(patient *domain.Patient) web.PatientResponse {
	return web.PatientResponse{
<<<<<<< Updated upstream
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======
>>>>>>> Stashed changes
		ID: 		  patient.ID,
=======
		ID: 	     patient.ID,
>>>>>>> Stashed changes
<<<<<<< Updated upstream
>>>>>>> Stashed changes
=======
>>>>>>> Stashed changes
		Name:         patient.Name,
		Email:        patient.Email,
		Nik:          patient.Nik,
		Birthday:     patient.Birthday,
		Age:          patient.Age,
		Address:      patient.Address,
		Gender:       patient.Gender,
		Phone_number: patient.Phone_number,
	}
}

func ConvertPatientResponse(patients []domain.Patient) []web.PatientResponse {
	var results []web.PatientResponse
	for _, patient := range patients {
		patientResponse := web.PatientResponse{
<<<<<<< Updated upstream
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======
>>>>>>> Stashed changes
			ID:           patient.ID,
=======
			ID: 		  patient.ID,
>>>>>>> Stashed changes
<<<<<<< Updated upstream
>>>>>>> Stashed changes
=======
>>>>>>> Stashed changes
			Name:         patient.Name,
			Email:        patient.Email,
			Nik:          patient.Nik,
			Birthday:     patient.Birthday,
			Age:          patient.Age,
			Address:      patient.Address,
			Gender:       patient.Gender,
			Phone_number: patient.Phone_number,
		}
		results = append(results, patientResponse)
	}
	return results
}