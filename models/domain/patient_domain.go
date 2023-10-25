package domain

type Patient struct {
	ID           uint
	Name         string
	Email        string
	Password     string
	Nik          string
<<<<<<< Updated upstream
	Birthday     string
=======
<<<<<<< Updated upstream
	Birthday      time.Time
>>>>>>> Stashed changes
	Age          uint64
=======
	Birthday     string
	Age          uint
>>>>>>> Stashed changes
	Address      string
	Gender       string
	Phone_number string
}