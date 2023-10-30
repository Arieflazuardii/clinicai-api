package web

type PatientLoginResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type PatientResponse struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Nik          string `json:"nik"`
	Birthday     string `json:"birthday"`
	Age          uint   `json:"age"`
	Address      string `json:"address"`
	Gender       string `json:"gender"`
	Phone_number string `json:"phone_number"`
}