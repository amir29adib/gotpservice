package utils

import (
	"os"
    "time"
    "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(userID uint, phone string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "phone":   phone,
        "exp":     time.Now().Add(24 * time.Hour).Unix(), // 24h expiry
        "iat":     time.Now().Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}
