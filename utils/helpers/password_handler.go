package helpers

import (
	"os"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	result, _ := bcrypt.GenerateFromPassword([]byte(password + os.Getenv("SALT")), bcrypt.DefaultCost)
	return string(result)
}

func ComparePassword(hash, password string) error {
	passwordWithSalt := []byte(password + os.Getenv("SALT"))
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwordWithSalt))
	if err != nil {
		return err
	}

	return nil
}