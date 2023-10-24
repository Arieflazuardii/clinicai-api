package request

import (
	"clinicai-api/models/domain"
	"clinicai-api/models/schema"
	"clinicai-api/models/web"
)


func PatientCreateRequestToPatientDomain(request web.PatientCreateRequest) *domain.Patient {
	return &domain.Patient{
		Name:         request.Name,
		Email:        request.Email,
		Password:     request.Password,
		Nik:          request.Nik,
		Birthday: request.Birthday,
		Age:          request.Age,
		Address:      request.Address,
		Gender:       request.Gender,
		Phone_number: request.Phone_number,
	}
}

func PatientLoginRequestToPatientDomain(request web.PatientLoginRequest) *domain.Patient {
	return &domain.Patient{
		Email:    request.Email,
		Password: request.Password,
	}
}

func PatientUpdateRequestToPatientDomain(request web.PatientUpdateRequest) *domain.Patient {
	return &domain.Patient{
		Name:         request.Name,
		Email:        request.Email,
		Password:     request.Password,
		Nik:          request.Nik,
		Birthday:     request.Birthday,
		Age:          request.Age,
		Address:      request.Address,
		Gender:       request.Gender,
		Phone_number: request.Phone_number,
	}
}

func PatientDomaintoPatientSchema(request domain.Patient) *schema.Patient{
	return &schema.Patient{
		Name:         request.Name,
		Email:        request.Email,
		Password:     request.Password,
		Nik:          request.Nik,      
		Birthday:     request.Birthday,
		Age:          request.Age,
		Address:      request.Address,
		Gender:       request.Gender,
		Phone_number: request.Phone_number,
	}
}