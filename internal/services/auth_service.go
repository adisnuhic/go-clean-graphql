package services

import (
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// IAuthService interface
type IAuthService interface {
	GeneratePasswordHash(password string) (string, error)
	ComparePasswordHash(password string, hash string) bool
	GenerateAccessToken(userID uint64, email string, roleID uint64, organisationID uint64) (string, error)
}

type authService struct {
	JWTSecret string
}

// NewAuthService -
func NewAuthService(jwtSecret string) IAuthService {
	return &authService{
		JWTSecret: jwtSecret,
	}
}

// GeneratePasswordHash generates hash for provided password
func (s authService) GeneratePasswordHash(password string) (string, error) {
	hashByte, errHashByte := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if errHashByte != nil {
		return "", errHashByte
	}

	return string(hashByte), nil
}

// GeneratePasswordHash generates hash for provided password
func (s authService) ComparePasswordHash(password string, hash string) bool {
	errPassword := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return errPassword == nil
}

// GenerateAccessToken generates access token
func (s authService) GenerateAccessToken(userID uint64, email string, roleID uint64, organisationID uint64) (string, error) {
	jwtSecret := []byte(s.JWTSecret)
	expirationTime := time.Now().UTC().Add(3000 * time.Minute)

	type CustomClaims struct {
		ID    uint64 `json:"id"`
		Email string `json:"email"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims := CustomClaims{
		userID,
		email,
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// JWT implementation
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
