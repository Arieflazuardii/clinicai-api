package response

import (
	"clinicai-api/models/domain"
	"clinicai-api/models/schema"
	"clinicai-api/models/web"
)

func DoctorDomainToDoctorLoginResponse(doctor *domain.Doctor) web.DoctorLoginResponse {
	return web.DoctorLoginResponse{
		Name:   doctor.Name,
		Email:  doctor.Email,
	}
}

func DoctorSchemaToDoctorDomain(doctor *schema.Doctor) *domain.Doctor {
	return &domain.Doctor{
<<<<<<< Updated upstream
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======
>>>>>>> Stashed changes
		ID:           doctor.ID,
=======
		ID: doctor.ID,
>>>>>>> Stashed changes
<<<<<<< Updated upstream
>>>>>>> Stashed changes
=======
>>>>>>> Stashed changes
		Name: 	      doctor.Name,
		Email:        doctor.Email,
		Password: 	  doctor.Password,
		Nik:          doctor.Nik,
		Specialization: doctor.Specialization,
		Gender:       doctor.Gender,
		Phone_number: doctor.Phone_number,
	}
}

func DoctorDomainToDoctorResponse(doctor *domain.Doctor) web.DoctorResponse {
	return web.DoctorResponse{
<<<<<<< Updated upstream
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======
>>>>>>> Stashed changes
		ID: 		  doctor.ID,
=======
		ID: doctor.ID,
>>>>>>> Stashed changes
<<<<<<< Updated upstream
>>>>>>> Stashed changes
=======
>>>>>>> Stashed changes
		Name:         doctor.Name,
		Email:        doctor.Email,
		Password:     doctor.Password,
		Nik:          doctor.Nik,
		Specialization: doctor.Specialization,
		Gender:       doctor.Gender,
		Phone_number: doctor.Phone_number,
	}
}

func ConvertDoctorResponse(doctors []domain.Doctor) []web.DoctorResponse {
	var results []web.DoctorResponse
	for _, doctor := range doctors {
		doctorResponse := web.DoctorResponse{
<<<<<<< Updated upstream
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======
>>>>>>> Stashed changes
			ID:          doctor.ID,
=======
			ID: doctor.ID,
>>>>>>> Stashed changes
<<<<<<< Updated upstream
>>>>>>> Stashed changes
=======
>>>>>>> Stashed changes
			Name:        doctor.Name,
			Email:       doctor.Email,
			Password:    doctor.Password,
			Nik:         doctor.Nik,
			Specialization: doctor.Specialization,
			Gender:      doctor.Gender,
			Phone_number:doctor.Phone_number,
		}
		results = append(results,doctorResponse)
	}
	return results
}