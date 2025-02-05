package signup

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "auditcue/internal/models"
    "auditcue/internal/signup/service"
)

// Handler for user signup
func SignupHandler(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    if err := service.CreateUser(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}