package service_test

import (
	mocks "ProjectFlour/internal/mocks"
	"ProjectFlour/internal/model"
	"ProjectFlour/internal/service"
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestCreateUser_Success(t *testing.T) {
	mockStrg := new(mocks.MockAuthorizationStorage)

	user := model.User{Name: "John", Username: "JohnMarston", Password: "qazwsxedc"}
	mockStrg.On("CreateUser", mock.AnythingOfType("model.User")).Return(1, nil)

	authService := service.NewAuthService(mockStrg)
	id, err := authService.CreateUser(user)

	assert.NoError(t, err)
	assert.Equal(t, 1, id)
	mockStrg.AssertExpectations(t)
}

func TestCreateUser_Error(t *testing.T) {
	mockStorage := new(mocks.MockAuthorizationStorage)

	user := model.User{Name: "John", Username: "john", Password: "pass123"}
	mockStorage.On("CreateUser", mock.AnythingOfType("model.User")).
		Return(0, errors.New("db error"))

	authService := service.NewAuthService(mockStorage)
	id, err := authService.CreateUser(user)

	assert.Error(t, err)
	assert.Equal(t, 0, id)
	mockStorage.AssertExpectations(t)
}

func TestGenerateToken_Success(t *testing.T) {
	mockStorage := new(mocks.MockAuthorizationStorage)

	mockStorage.On("GetUser", "john", mock.AnythingOfType("string")).
		Return(model.User{ID: 1, Username: "john"}, nil)

	authService := service.NewAuthService(mockStorage)
	token, err := authService.GenerateToken("john", "pass123")

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	mockStorage.AssertExpectations(t)
}

func TestGenerateToken_UserNotFound(t *testing.T) {
	mockStorage := new(mocks.MockAuthorizationStorage)

	mockStorage.On("GetUser", "john", mock.AnythingOfType("string")).
		Return(model.User{}, errors.New("user not found"))

	authService := service.NewAuthService(mockStorage)
	token, err := authService.GenerateToken("john", "pass123")

	assert.Error(t, err)
	assert.Empty(t, token)
	assert.Contains(t, err.Error(), "user not found")
	mockStorage.AssertExpectations(t)
}

func TestParseToken_Success(t *testing.T) {
	// создаем тестовый токен
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
		IssuedAt:  time.Now().Unix(),
	}
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  42,
		"username": "john",
		"exp":      claims.ExpiresAt,
		"iat":      claims.IssuedAt,
	})
	tokenStr, _ := tokenObj.SignedString([]byte("qrkjk#4#%35FSFJlja#4353KSFjH"))

	authService := service.NewAuthService(nil) // storage не нужен для ParseToken
	userID, err := authService.ParseToken(tokenStr)

	assert.NoError(t, err)
	assert.Equal(t, 42, userID)
}

func TestParseToken_Invalid(t *testing.T) {
	authService := service.NewAuthService(nil)

	_, err := authService.ParseToken("invalid.token.here")
	assert.Error(t, err)
}

// NOTE:
//	The password wasn't inspected, because GenerateToken() hashed it yourself
