package web

type PatientLoginResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type PatientResponse struct {
	ID           uint64 `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Nik          string `json:"nik"`
	Birthday     string `json:"birthday"`
	Age          uint64 `json:"age"`
	Address      string `json:"address"`
	Gender       string `json:"gender"`
	Phone_number string `json:"phone_number"`

}