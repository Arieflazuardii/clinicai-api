package domain

type Patient struct {
	ID           uint
	Name         string
	Email        string
	Password     string
	Nik          string
	Birthday     string
	Age          uint
	Address      string
	Gender       string
	Phone_number string
	Registration []Registration
}