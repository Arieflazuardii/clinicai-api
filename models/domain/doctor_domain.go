package domain

type Doctor struct {
	ID             uint
	Name           string
	Email          string
	Password       string
	Nik            string
	Specialization string
	Gender         string
	Phone_number   string
	Schedule       []Schedule
}