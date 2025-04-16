package auth

import (
	"errors"
	"time"

	"user_service/internal/domain"

	"github.com/golang-jwt/jwt/v5"
)

// TokenService handles JWT token generation and validation
type TokenService struct {
	jwtSecret string
}

// Claims represents JWT claims
type Claims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// NewTokenService creates a new TokenService instance
func NewTokenService(jwtSecret string) *TokenService {
	return &TokenService{
		jwtSecret: jwtSecret,
	}
}

// GenerateToken creates a new JWT token for a user
func (s *TokenService) GenerateToken(user *domain.User) (string, int64, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	expireAt := expirationTime.Unix()

	claims := &Claims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", 0, err
	}

	return tokenString, expireAt, nil
}

// ValidateToken validates a JWT token and returns the claims
func (s *TokenService) ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
