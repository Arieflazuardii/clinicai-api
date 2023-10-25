package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


func GenerateTokenPatient(PatientID uint) (string, error) {
	jwtSecret := []byte(os.Getenv("SECRET_KEY"))

	claims := jwt.MapClaims{
		"sub": PatientID,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
		"iat": time.Now().Unix(),
		"role": "Patient",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateTokenDoctor(DoctorID uint) (string, error) {
    jwtSecret := []byte(os.Getenv("SECRET_KEY"))

    claims := jwt.MapClaims{
        "sub":   DoctorID,
        "exp":   time.Now().Add(time.Hour * 1).Unix(),
        "iat":   time.Now().Unix(),
		"role": "Doctor",
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}