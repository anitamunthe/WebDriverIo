package handlers

import (
	"net/http"

	"Rest-Api/db"
	"Rest-Api/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var input models.Akun
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if input.Username == "" && input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username dan password harus diisi"})
		return
	} else if input.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username harus diisi"})
		return
	} else if input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password harus diisi"})
		return
	}

	var akun models.Akun
	if err := db.DB.Table("akun").Where("username = ?", input.Username).First(&akun).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(akun.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": akun})
}

