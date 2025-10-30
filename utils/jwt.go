package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type JwtCustomClaims struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(id int64, email string) (string, error) {
	// Set custom claims
	customClaims := &JwtCustomClaims{
		ID:    uint(id),
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)), // Token expires in 12 hours
		},
	}

	// Create token with customClaims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)

	// Generate encoded token and send it as response.
	jwtToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func GetSecretKey() []byte {
	return secretKey
}
