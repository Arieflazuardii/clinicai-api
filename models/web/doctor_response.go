package web

type DoctorLoginResponse struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Specialization string `json:"specialization"`
	Token          string `json:"token"`
}

type DoctorResponse struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Nik            string `json:"nik"`
	Specialization string `json:"specialization"`
	Gender         string `json:"gender"`
	Phone_number   string `json:"phone_number"`
}