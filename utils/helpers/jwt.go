package helpers

import (
	"clinicai-api/models/web"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(PatientLoginResponse *web.PatientLoginResponse, id int) (string, error) {
	expireTime := time.Now().Add(time.Hour * 168).Unix()
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["name"] = PatientLoginResponse.Name
	claims["email"] = PatientLoginResponse.Email
	claims["exp"] = expireTime

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return validToken, nil
}