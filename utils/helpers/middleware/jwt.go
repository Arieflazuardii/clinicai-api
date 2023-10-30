package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)


func GenerateTokenPatient(PatientID uint) (string, error) {
	jwtSecret := []byte(os.Getenv("SECRET_KEY"))

	claims := jwt.MapClaims{
		"id": PatientID,
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
        "id":   DoctorID,
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

func ExtractTokenPatientId(e echo.Context) float64 {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		PatientId := claims["id"].(float64)
		return PatientId
	}
	return 0
}


func ExtractTokenDoctorId(e echo.Context) float64 {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
	DoctorId := claims["id"].(float64)
		return DoctorId
	}
	return 0

}