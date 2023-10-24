package web

type DoctorCreateRequest struct {
	Name           string `json:"name" validate:"required,min=1,max=255"`
	Email          string `json:"email" validate:"required,email,min=1,max=255"`
	Password       string `json:"password" validate:"required,min=8,max=255"`
	Nik            string `json:"nik" validate:"required,min=16,max=16"`
	Specialization string `json:"specialization"`
	Gender         string `json:"gender"`
	Phone_number   string `json:"phone_number"`
}

type DoctorLoginRequest struct {
	Email    string `json:"email" validate:"required,email,min=1,max=255"`
	Password string `json:"password" validate:"required,max=255"`
}

type DoctorUpdateRequest struct {
	Name           string `json:"name" validate:"required,min=1,max=255"`
	Email          string `json:"email" validate:"required,email,min=1,max=255"`
	Password       string `json:"password" validate:"required,min=8,max=255"`
	Nik            string `json:"nik" validate:"required,min=16,max=16"`
	Specialization string `json:"specialization"`
	Gender         string `json:"gender"`
	Phone_number   string `json:"phone_number" validate:"required,min=10,max=16"`
}