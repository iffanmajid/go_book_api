package auth

import (
	"os"
	"time"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type authUseCase struct {
	jwtSecret []byte
}

func NewAuthUseCase() UseCase {
	return &authUseCase{
		jwtSecret: []byte(os.Getenv("SECRET_TOKEN")),
	}
}

func (u *authUseCase) GenerateToken(username, password string) (string, error) {
	// In a real application, you would validate credentials against a database
	if username != "admin" || password != "password" {
		return "", errors.New("invalid credentials")
	}

	expirationTime := time.Now().Add(15 * time.Minute)
	claims := jwt.MapClaims{
		"username": username,
		"exp":      expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(u.jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (u *authUseCase) ValidateToken(tokenString string) error {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return u.jwtSecret, nil
	})

	return err
}