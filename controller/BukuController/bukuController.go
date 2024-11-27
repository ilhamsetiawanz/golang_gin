package Bukucontroller

import (
	"golang_gin/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Read(c *gin.Context) {
	var buku []database.Buku

	database.DB.Find(&buku)

	if buku == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Buku Kosong",
		})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"Buku": buku})
	}
}

func ReadById(c *gin.Context) {
	var buku database.Buku
	id := c.Param("id")

	if err := database.DB.First(&buku, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Buku tidak ditemukan",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})
		}

	}
	c.JSON(http.StatusOK, gin.H{
		"Buku": buku,
	})

}

func Create(c *gin.Context) {
	var buku database.Buku
	if err := c.ShouldBindJSON(&buku); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&buku)

	c.JSON(http.StatusCreated, gin.H{
		"buku": buku,
	})
}

func Update(c *gin.Context) {
	var buku database.Buku
	id := c.Param("id")

	// Bind JSON ke struct, cek apakah valid
	if err := c.ShouldBindJSON(&buku); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid input",
			"detail": err.Error(),
		})
		return
	}

	// Update record di database
	if database.DB.Model(&buku).Where("id = ?", id).Updates(&buku).RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Buku tidak ditemukan",
		})
		return
	}

	// Jika berhasil diperbarui
	c.JSON(http.StatusAccepted, gin.H{
		"message": "Buku Berhasil Diperbarui",
		"buku":    buku,
	})
}

func Delete(c *gin.Context) {
	var buku database.Buku

	id := c.Param("id")

	if database.DB.Model(&buku).Where("id = ?", id).Delete(&buku).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Buku tidak ditemukan",
		})
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Buku Berhasil Dihapus",
	})

}
