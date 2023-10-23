package domain

import "time"


type Patient struct {
	ID           uint64
	Name         string
	Email        string
	Password     string
	Nik          string
	Birthday      time.Time
	Age          uint64
	Address      string
	Gender       string
	Phone_number string
}