package auth

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// Handler struct to hold dependencies
type Handler struct {
    service *Service
}

// NewHandler creates a new Handler for authentication
func NewHandler(service *Service) *Handler {
    return &Handler{service: service}
}

// Login handles user login requests
func (h *Handler) Login(c *gin.Context) {
    var loginRequest LoginRequest
    if err := c.ShouldBindJSON(&loginRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    token, err := h.service.Login(loginRequest.Email, loginRequest.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

// Logout handles user logout requests
func (h *Handler) Logout(c *gin.Context) {
    // Logic for logging out the user (e.g., invalidating the token)
    c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

// LoginRequest represents the structure of a login request
type LoginRequest struct {
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}