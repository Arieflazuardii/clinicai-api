package web

type PatientLoginResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type PatientResponse struct {
<<<<<<< Updated upstream
<<<<<<< Updated upstream
	ID           uint64 `json:"id"`
=======
<<<<<<< Updated upstream
	ID			uint64     `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Nik          string    `json:"nik"`
	Birthday      time.Time `json:"birthday"`
	Age          uint64    `json:"age"`
	Address      string    `json:"address"`
	Gender       string    `json:"gender"`
	Phone_number string    `json:"phone_number"`
=======
	ID           uint `json:"id"`
>>>>>>> Stashed changes
=======
	ID           uint   `json:"id"`
>>>>>>> Stashed changes
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Nik          string `json:"nik"`
	Birthday     string `json:"birthday"`
<<<<<<< Updated upstream
	Age          uint64 `json:"age"`
=======
	Age          uint `json:"age"`
>>>>>>> Stashed changes
	Address      string `json:"address"`
	Gender       string `json:"gender"`
	Phone_number string `json:"phone_number"`
<<<<<<< Updated upstream

<<<<<<< Updated upstream
=======
>>>>>>> Stashed changes
>>>>>>> Stashed changes
=======
>>>>>>> Stashed changes
}