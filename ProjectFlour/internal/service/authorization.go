package service

import (
	"ProjectFlour/internal/model"
	"ProjectFlour/internal/storage"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 24 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
}

type AuthorizationService struct {
	storage storage.AuthorizationStorage
}

func NewAuthService(storage storage.AuthorizationStorage) *AuthorizationService {
	return &AuthorizationService{storage: storage}
}

func (s *AuthorizationService) CreateUser(user model.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.storage.CreateUser(user)
}

func (s *AuthorizationService) GenerateToken(username string, password string) (string, error) {
	hashedPassword := generatePasswordHash(password)

	user, err := s.storage.GetUser(username, hashedPassword)
	if err != nil {
		return "", fmt.Errorf("authentication failed: %w", err)
	}

	if user.ID == 0 {
		return "", errors.New("user not found")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		UserID:   user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthorizationService) ParseToken(accessToken string) (int, error) {
	const op = "service.authorization.ParseToken"

	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not valid")
	}

	return claims.UserID, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
