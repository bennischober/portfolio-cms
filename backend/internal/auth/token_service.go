package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type ITokenService interface {
	GenerateToken(username string, expireTime time.Duration) (string, error)
	ValidateToken(tokenString string) (*Claims, error)
}

type TokenService struct {
	JwtKey []byte
}

func NewTokenService() *TokenService {
    return &TokenService{
        JwtKey: []byte(os.Getenv("JWT_SECRET")),
    }
}

func (ts *TokenService) GenerateToken(username string, expireTime time.Duration) (string, error) {
	expirationTime := time.Now().Add(expireTime)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(ts.JwtKey) // Sign and get the complete encoded token as a string using the secret

	return tokenString, err
}

func (ts *TokenService) ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return ts.JwtKey, nil
	})

	if !token.Valid {
		return nil, err
	}

	return claims, err
}
