package auth

import (
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAuthUseCase_GenerateToken(t *testing.T) {
	// Set test secret
	os.Setenv("SECRET_TOKEN", "test-secret")
	useCase := NewAuthUseCase()

	// Test valid credentials
	token, err := useCase.GenerateToken("admin", "password")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Test invalid credentials
	token, err = useCase.GenerateToken("wrong", "credentials")
	assert.Error(t, err)
	assert.Empty(t, token)
}

func TestAuthUseCase_ValidateToken(t *testing.T) {
	// Set test secret
	os.Setenv("SECRET_TOKEN", "test-secret")
	useCase := NewAuthUseCase()

	// Generate a valid token first
	token, err := useCase.GenerateToken("admin", "password")
	assert.NoError(t, err)

	// Test valid token
	err = useCase.ValidateToken(token)
	assert.NoError(t, err)

	// Test invalid token
	err = useCase.ValidateToken("invalid-token")
	assert.Error(t, err)
}