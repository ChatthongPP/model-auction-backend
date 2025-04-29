package auth

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	jwtSecret []byte
	once      sync.Once
)

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	RoleID string `json:"role_id"`
	jwt.RegisteredClaims
}

func loadSecret() {
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
}

func GenerateJWT(userID uint, email string, roleID uint) (string, error) {
	once.Do(loadSecret)

	claims := &Claims{
		UserID: fmt.Sprintf("%d", userID),
		Email:  email,
		RoleID: fmt.Sprintf("%d", roleID),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenStr string) (*Claims, error) {
	once.Do(loadSecret)

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("could not parse claims")
	}

	return claims, nil
}
