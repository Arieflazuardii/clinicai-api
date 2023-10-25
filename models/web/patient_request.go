package web

<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
import "time"

>>>>>>> Stashed changes

type PatientCreateRequest struct {
	Name         string    `json:"name" validate:"required,min=1,max=255"`
	Email        string    `json:"email" validate:"required,email,min=1,max=255"`
	Password     string    `json:"password" validate:"required,min=8,max=255"`
	Nik          string    `json:"nik" validate:"required,min=16,max=16"`
	Birthday     string `json:"birthday"`
	Age          uint64    `json:"age"`
	Address      string    `json:"address"`
	Gender       string    `json:"gender"`
	Phone_number string    `json:"phone_number"`
=======
type PatientCreateRequest struct {
	Name         string `json:"name" validate:"required,min=1,max=255"`
	Email        string `json:"email" validate:"required,email,min=1,max=255"`
	Password     string `json:"password" validate:"required,min=8,max=255"`
	Nik          string `json:"nik" validate:"required,min=16,max=16"`
	Birthday     string `json:"birthday"`
	Age          uint   `json:"age"`
	Address      string `json:"address"`
	Gender       string `json:"gender"`
	Phone_number string `json:"phone_number"`
>>>>>>> Stashed changes
}

type PatientLoginRequest struct {
	Email    string `json:"email" validate:"required,email,min=1,max=255"`
	Password string `json:"password" validate:"required,max=255"`
}

type PatientUpdateRequest struct {
<<<<<<< Updated upstream
	Name         string    `json:"name" validate:"required,min=1,max=255"`
	Email        string    `json:"email" validate:"required,email,min=1,max=255"`
	Password     string    `json:"password" validate:"required,min=8,max=255"`
	Nik          string    `json:"nik" validate:"required,min=16,max=16"`
	Birthday     string `json:"birthday"`
	Age          uint64    `json:"age"`
	Address      string    `json:"address"`
	Gender       string    `json:"gender"`
	Phone_number string    `json:"phone_number" validate:"required,min=10,max=16"`
=======
	Name         string `json:"name" validate:"required,min=1,max=255"`
	Email        string `json:"email" validate:"required,email,min=1,max=255"`
	Password     string `json:"password" validate:"required,min=8,max=255"`
	Nik          string `json:"nik" validate:"required,min=16,max=16"`
	Birthday     string `json:"birthday"`
	Age          uint   `json:"age"`
	Address      string `json:"address"`
	Gender       string `json:"gender"`
	Phone_number string `json:"phone_number" validate:"required,min=10,max=16"`
>>>>>>> Stashed changes
}