package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go_book_api/domain"
	"go_book_api/usecase/auth"
)

type AuthHandler struct {
	authUseCase auth.UseCase
}

func NewAuthHandler(router *gin.Engine, useCase auth.UseCase) {
	handler := &AuthHandler{authUseCase: useCase}

	router.POST("/token", handler.GenerateToken)
}

func (h *AuthHandler) GenerateToken(c *gin.Context) {
	var loginReq domain.LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authUseCase.GenerateToken(loginReq.Username, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.LoginResponse{Token: token})
}