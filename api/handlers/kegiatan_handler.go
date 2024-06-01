// api/handlers/kegiatan_handler.go

package handlers

import (
	"Rest-Api/db"
	"Rest-Api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllKegiatan(c *gin.Context) {
	var kegiatans []models.Kegiatan
	db.DB.Find(&kegiatans)

	c.JSON(http.StatusOK, kegiatans)
}

func CreateKegiatan(c *gin.Context) {
	var newKegiatan models.Kegiatan
	if err := c.ShouldBindJSON(&newKegiatan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Validasi setiap field yang wajib diisi
	if newKegiatan.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title harus diisi"})
		return
	}
	if newKegiatan.Start.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Start harus diisi"})
		return
	}
	if newKegiatan.End.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "End harus diisi"})
		return
	}
	if newKegiatan.Color == "" || len(newKegiatan.Color) != 7 || newKegiatan.Color[0] != '#' {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Color harus diisi dengan format hex (#RRGGBB)"})
		return
	}
	if newKegiatan.Status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status harus diisi"})
		return
	}
	if newKegiatan.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Description harus diisi"})
		return
	}
	currentTime := time.Now()
	newKegiatan.CreatedAt = currentTime
	newKegiatan.UpdatedAt = currentTime

	db.DB.Create(&newKegiatan)
	c.JSON(http.StatusOK, newKegiatan)
}

func UpdateKegiatan(c *gin.Context) {
	kegiatanID := c.Param("id")
	var updatedKegiatan models.Kegiatan

	if err := db.DB.Where("id = ?", kegiatanID).First(&updatedKegiatan).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kegiatan not found"})
		return
	}

	if err := c.ShouldBindJSON(&updatedKegiatan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}
	updatedKegiatan.UpdatedAt = time.Now()

	db.DB.Save(&updatedKegiatan)
	c.JSON(http.StatusOK, updatedKegiatan)
}
