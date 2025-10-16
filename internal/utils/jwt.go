package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Replace with your secret key (keep it private!)
var jwtSecret = []byte("your_super_secret_key")

func CreateJwt(userId int) (string, error) {
	// Implementation for creating a JWT token
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // expires in 1 day
		"iat":     time.Now().Unix(),                     // issued at
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// VerifyJWT parses and validates a token, returning the user ID if valid
func VerifyJWT(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		// Ensure the signing method is HMAC
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})
	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	// Extract user ID
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if userIDFloat, ok := claims["user_id"].(float64); ok {
			return int(userIDFloat), nil
		}
		return 0, errors.New("user_id missing or invalid type")
	}

	return 0, errors.New("invalid claims")
}
