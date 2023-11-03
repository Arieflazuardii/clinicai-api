package request

import (
	"clinicai-api/models/domain"
	"clinicai-api/models/schema"
	"clinicai-api/models/web"
)

func RegistrationDomainToRegistrationSchema(request domain.Registration) *schema.Registration {
	return &schema.Registration{
		PatientID: request.PatientID,
		DoctorID: request.DoctorID,
		ScheduleID: request.ScheduleID,
		Complaint: request.Complaint,
	}
}

func RegistrationCreateRequestToRegistrationDomain(request web.RegistrationCreateRequest) *domain.Registration {
	return &domain.Registration{
		PatientID: request.PatientID,
		DoctorID: request.DoctorID,
		ScheduleID: request.ScheduleID,
		Complaint: request.Complaint,
	}
}

func RegistrationUpdateRequestToRegistrationDomain(request web.RegistrationUpdateRequest) *domain.Registration {
	return &domain.Registration{
		PatientID: request.PatientID,
		DoctorID: request.DoctorID,
		ScheduleID: request.ScheduleID,
		Complaint: request.Complaint,
	}
}
