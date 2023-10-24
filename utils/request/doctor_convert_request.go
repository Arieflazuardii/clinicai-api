package request

import (
	"clinicai-api/models/domain"
	"clinicai-api/models/schema"
	"clinicai-api/models/web"
)


func DoctorCreateRequestToDoctorDomain(request web.DoctorCreateRequest) *domain.Doctor {
	return &domain.Doctor{
		Name:         		request.Name,
		Email:       		request.Email,
		Password:     		request.Password,
		Nik:          		request.Nik,      
		Specialization:     request.Specialization,
		Gender:       		request.Gender,
		Phone_number: 		request.Phone_number,
	}
}

func DoctorLoginRequestToDoctorDomain(request web.DoctorLoginRequest) *domain.Doctor {
	return &domain.Doctor{
		Email:    			request.Email,
		Password: 			request.Password,
	}
}

func DoctorUpdateRequestToDoctorDomain(request web.DoctorUpdateRequest) *domain.Doctor {
	return &domain.Doctor{
		Name:         		request.Name,
		Email:        		request.Email,
		Password:     		request.Password,
		Nik:          		request.Nik,
		Specialization:     request.Specialization,
		Gender:       		request.Gender,
		Phone_number: 		request.Phone_number,
	}
}

func DoctorDomaintoDoctorSchema(request domain.Doctor) *schema.Doctor{
	return &schema.Doctor{
		Name:         request.Name,
		Email:        request.Email,
		Password:     request.Password,
		Nik:          request.Nik,      
		Specialization:     request.Specialization,
		Gender:       request.Gender,
		Phone_number: request.Phone_number,
	}
}