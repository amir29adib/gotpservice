package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
    UserID uint   `json:"user_id"`
    Phone  string `json:"phone"`
    jwt.RegisteredClaims
}

func GenerateJWT(userID uint, phone string) (string, error) {
    claims := Claims{
        UserID: userID,
        Phone:  phone,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

func ParseJWT(tokenString string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte("secret-key"), nil
    })

    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
        return claims, nil
    } else {
        return nil, err
    }
}
