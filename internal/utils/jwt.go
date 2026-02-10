package utils

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte

func init() {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("jwt secret is not set")
	}
	jwtSecret = []byte(secret)
}

// claims struct
type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func CreateJwt(userId int) (string, error) {
	// Implementation for creating a JWT token
	claims := Claims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func VerifyJWT(tokenString string) (int, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
		if t.Method != jwt.SigningMethodHS256 {
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

	return claims.UserID, nil
}
