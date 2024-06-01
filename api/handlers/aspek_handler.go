// api/handlers/aspek_handler.go

package handlers

import (
	"Rest-Api/db"
	"Rest-Api/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateAspek(c *gin.Context) {
    var newAspek models.Aspek
    if err := c.ShouldBindJSON(&newAspek); err != nil {
        // Log kesalahan dalam binding data
        fmt.Println("Error binding JSON:", err.Error())

        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
        return
    }

    // Cek setiap field yang wajib diisi
    if newAspek.NamaAspek == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Nama aspek harus diisi"})
        return
    }
    if newAspek.Kode == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Kode harus diisi"})
        return
    }
    if newAspek.KelasID == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Kelas ID harus diisi"})
        return
    }

    // Lanjutkan proses simpan ke database jika semua field wajib diisi
    currentTime := time.Now()
    newAspek.CreatedAt = currentTime
    newAspek.UpdatedAt = currentTime

    db.DB.Create(&newAspek)
    c.JSON(http.StatusOK, newAspek)
}

func GetAspekList(c *gin.Context) {
	var aspekList []models.Aspek
	db.DB.Find(&aspekList)
	c.JSON(http.StatusOK, aspekList)
}

func GetAspekByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var aspek models.Aspek
	if err := db.DB.Where("id = ?", id).First(&aspek).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, aspek)
}

func UpdateAspek(c *gin.Context) {
	id := c.Params.ByName("id")
	var aspek models.Aspek
	if err := db.DB.Where("id = ?", id).First(&aspek).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var updatedAspek models.Aspek
	if err := c.BindJSON(&updatedAspek); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	currentTime := time.Now()
	updatedAspek.UpdatedAt = currentTime

	db.DB.Model(&aspek).Updates(updatedAspek)
	c.JSON(http.StatusOK, updatedAspek)
}

func DeleteAspek(c *gin.Context) {
	id := c.Params.ByName("id")
	var aspek models.Aspek
	if err := db.DB.Where("id = ?", id).First(&aspek).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	db.DB.Delete(&aspek)
	c.JSON(http.StatusOK, gin.H{"message": "Aspek deleted"})
}
