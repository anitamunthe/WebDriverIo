// api/handlers/kelas_handler.go

package handlers

import (
	"Rest-Api/db"
	"Rest-Api/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetKelasList handles the GET /kelas endpoint
func GetKelasList(c *gin.Context) {
	var kelas []models.Kelas
	db.DB.Find(&kelas)

	c.JSON(200, kelas)
}

// GetKelasByID handles the GET /kelas/:id endpoint
func GetKelasByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var kelas models.Kelas
	if err := db.DB.Where("id = ?", id).First(&kelas).Error; err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, kelas)
	}
}

func CreateKelas(c *gin.Context) {
	var newKelas models.Kelas
	if err := c.ShouldBindJSON(&newKelas); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Validasi setiap field yang wajib diisi
	if newKelas.NamaKelas == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama kelas harus diisi"})
		return
	}
	if newKelas.Kode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Kode kelas harus diisi"})
		return
	}

	currentTime := time.Now()
	newKelas.CreatedAt = currentTime
	newKelas.UpdatedAt = currentTime

	db.DB.Create(&newKelas)
	c.JSON(http.StatusOK, newKelas)
}

// UpdateKelas handles the PUT /kelas/:id endpoint
func UpdateKelas(c *gin.Context) {
	id := c.Params.ByName("id")
	var kelas models.Kelas
	if err := db.DB.Where("id = ?", id).First(&kelas).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.BindJSON(&kelas)
	db.DB.Save(&kelas)
	c.JSON(200, kelas)
}

// DeleteKelas handles the DELETE /kelas/:id endpoint
func DeleteKelas(c *gin.Context) {
	id := c.Params.ByName("id")
	var kelas models.Kelas
	d := db.DB.Where("id = ?", id).Delete(&kelas)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
