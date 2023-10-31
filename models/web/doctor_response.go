package web

type DoctorLoginResponse struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Specialization string `json:"specialization"`
	Token          string `json:"token"`
}

type DoctorResponse struct {
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Nik             string `json:"nik"`
	Specialization string `json:"specialization"`
	Gender          string `json:"gender"`
	Phone_number    string `json:"phone_number"`
=======
	ID             uint   `json:"id"`
>>>>>>> Stashed changes
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Nik            string `json:"nik"`
	Specialization string `json:"specialization"`
	Gender         string `json:"gender"`
	Phone_number   string `json:"phone_number"`
<<<<<<< Updated upstream
=======
>>>>>>> Stashed changes
>>>>>>> Stashed changes
}