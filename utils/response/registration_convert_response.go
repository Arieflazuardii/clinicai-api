package response

import (
	"clinicai-api/models/domain"
	"clinicai-api/models/schema"
	"clinicai-api/models/web"
)

func RegistrationSchemaToRegistrationDomain(registration *schema.Registration) *domain.Registration {
	return &domain.Registration{
		ID: registration.ID,
		DoctorID: registration.DoctorID,
		PatientID: registration.PatientID,
		ScheduleID: registration.ScheduleID,
		Complaint: registration.Complaint,
	}
}

func RegistrationDomainToRegistrationResponse(registration *domain.Registration) web.RegistrationResponse {
	return web.RegistrationResponse{
		ID: registration.ID,
		PatientName: registration.PatientName,
		DoctorName: registration.DoctorName,
		ScheduleDate: registration.ScheduleDate,
		Complaint: registration.Complaint,
	}
}

func CreateRegistrationDomainToRegistrationResponse(registration *domain.Registration) web.CreateRegistrationResponse{
	return web.CreateRegistrationResponse{
<<<<<<< Updated upstream
<<<<<<< Updated upstream
=======
		ID: registration.ID,
>>>>>>> Stashed changes
=======
		ID: registration.ID,
>>>>>>> Stashed changes
		PatientID: registration.PatientID,
		DoctorID: registration.DoctorID,
		ScheduleID: registration.ScheduleID,
		Complaint: registration.Complaint,
	}
}

func UpdateRegistrationDomainToRegistrationResponse(registration *domain.Registration) web.CreateRegistrationResponse{
	return web.CreateRegistrationResponse{
<<<<<<< Updated upstream
<<<<<<< Updated upstream
=======
		ID: registration.ID,
>>>>>>> Stashed changes
=======
		ID: registration.ID,
>>>>>>> Stashed changes
		PatientID: registration.PatientID,
		DoctorID: registration.DoctorID,
		ScheduleID: registration.ScheduleID,
		Complaint: registration.Complaint,

	}
}

func ConvertRegistrationResponse(registrations []domain.Registration) []web.RegistrationResponse {
	var results []web.RegistrationResponse
	for _, registration := range registrations {
		registrationResponse := web.RegistrationResponse{
<<<<<<< Updated upstream
<<<<<<< Updated upstream
=======
			ID: registration.ID,
>>>>>>> Stashed changes
=======
			ID: registration.ID,
>>>>>>> Stashed changes
			PatientName: registration.PatientName,
			DoctorName: registration.DoctorName,
			ScheduleDate: registration.ScheduleDate,
			Complaint: registration.Complaint,
		}
		results = append(results, registrationResponse)
		
	} 
	return results
}