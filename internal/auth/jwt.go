package auth

// JWT authentication utilities

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("change-this-secret")

func GenetateToken(userID uint) (string, error) {

	claims := jwt.MapClaims{

		"user_id": userID,
		"role":    "agent",
		"exp":     time.Now().Add(24 * 31 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
