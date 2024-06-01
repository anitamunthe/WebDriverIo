// api/handlers/poin_aspek_handler.go

package handlers

import (
	"Rest-Api/db"
	"Rest-Api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllPoinAspek untuk mendapatkan semua poin_aspek
func GetAllPoinAspek(c *gin.Context) {
    var poinAspeks []models.PoinAspek
    db.DB.Preload("Aspek").Find(&poinAspeks)

    c.JSON(http.StatusOK, poinAspeks)
}

// GetPoinAspekByID untuk mendapatkan poin_aspek berdasarkan ID
func GetPoinAspekByID(c *gin.Context) {
    var poinAspek models.PoinAspek
    id := c.Param("id")

    if err := db.DB.Preload("Aspek").Where("id = ?", id).First(&poinAspek).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "PoinAspek not found"})
        return
    }

    c.JSON(http.StatusOK, poinAspek)
}

// CreatePoinAspek untuk membuat poin_aspek baru
func CreatePoinAspek(c *gin.Context) {
	var newPoinAspek models.PoinAspek
	if err := c.ShouldBindJSON(&newPoinAspek); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Validasi setiap field yang wajib diisi
	if newPoinAspek.NamaPoin == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama point harus diisi"})
		return
	}
	if newPoinAspek.AspekID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Aspek ID harus diisi"})
		return
	}

	// Lakukan validasi tambahan jika diperlukan

	db.DB.Create(&newPoinAspek)
	c.JSON(http.StatusOK, newPoinAspek)
}

// UpdatePoinAspek untuk memperbarui poin_aspek yang ada
func UpdatePoinAspek(c *gin.Context) {
    var poinAspek models.PoinAspek
    id := c.Param("id")

    if err := db.DB.Where("id = ?", id).First(&poinAspek).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "PoinAspek not found"})
        return
    }

    var updatedPoinAspek models.PoinAspek
    if err := c.ShouldBindJSON(&updatedPoinAspek); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
        return
    }

    // Pastikan aspek_id yang dimasukkan ada dalam tabel aspek
    var aspek models.Aspek
    if err := db.DB.First(&aspek, updatedPoinAspek.AspekID).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid aspek_id"})
        return
    }

    db.DB.Model(&poinAspek).Updates(updatedPoinAspek)
    c.JSON(http.StatusOK, updatedPoinAspek)
}

// DeletePoinAspek untuk menghapus poin_aspek
func DeletePoinAspek(c *gin.Context) {
    var poinAspek models.PoinAspek
    id := c.Param("id")

    if err := db.DB.Where("id = ?", id).First(&poinAspek).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "PoinAspek not found"})
        return
    }

    db.DB.Delete(&poinAspek)
    c.JSON(http.StatusOK, gin.H{"message": "PoinAspek deleted"})
}
