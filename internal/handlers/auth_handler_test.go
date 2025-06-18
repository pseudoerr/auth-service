package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/pseudoerr/auth-service/internal/handlers"
	"github.com/pseudoerr/auth-service/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthHandler_Register(t *testing.T) {
	// Setup
	mockAuthService := &mocks.AuthService{}
	handler := handlers.NewAuthHandler(mockAuthService)

	// Test data
	registerReq := models.RegisterRequest{
		Email:    "test@example.com",
		Username: "testuser",
		Password: "TestPassword123",
	}

	authResponse := &models.AuthResponse{
		AccessToken:  "access-token",
		RefreshToken: "refresh-token",
		User: models.User{
			ID:       1,
			Email:    registerReq.Email,
			Username: registerReq.Username,
		},
	}

	// Mock expectations
	mockAuthService.On("Register", mock.AnythingOfType("*models.RegisterRequest")).Return(authResponse, nil)

	// Create request
	reqBody, _ := json.Marshal(registerReq)
	req := httptest.NewRequest("POST", "/auth/register", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	rr := httptest.NewRecorder()

	// Execute
	handler.Register(rr, req)

	// Assert
	assert.Equal(t, http.StatusCreated, rr.Code)

	var response models.AuthResponse
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, authResponse.AccessToken, response.AccessToken)
	assert.Equal(t, authResponse.User.Email, response.User.Email)

	mockAuthService.AssertExpectations(t)
}

func TestAuthHandler_Login(t *testing.T) {
	// Setup
	mockAuthService := &mocks.AuthService{}
	handler := handlers.NewAuthHandler(mockAuthService)

	// Test data
	loginReq := models.LoginRequest{
		Email:    "test@example.com",
		Password: "TestPassword123",
	}

	authResponse := &models.AuthResponse{
		AccessToken:  "access-token",
		RefreshToken: "refresh-token",
		User: models.User{
			ID:       1,
			Email:    loginReq.Email,
			Username: "testuser",
		},
	}

	// Mock expectations
	mockAuthService.On("Login", mock.AnythingOfType("*models.LoginRequest")).Return(authResponse, nil)

	// Create request
	reqBody, _ := json.Marshal(loginReq)
	req := httptest.NewRequest("POST", "/auth/login", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	rr := httptest.NewRecorder()

	// Execute
	handler.Login(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)

	var response models.AuthResponse
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, authResponse.AccessToken, response.AccessToken)

	mockAuthService.AssertExpectations(t)
}
