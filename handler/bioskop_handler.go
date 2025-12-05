package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"Tugas-13-Gin-dan-Postgres/database"
	"Tugas-13-Gin-dan-Postgres/model"
	"strconv"
)

type InputBioskop struct {
	Nama   string  `json:"nama"`
	Lokasi string  `json:"lokasi"`
	Rating float32 `json:"rating"`
}

func CreateBioskop(c *gin.Context) {
	nama := c.PostForm("nama")
	lokasi := c.PostForm("lokasi")
	ratingStr := c.PostForm("rating")

	if nama == "" || lokasi == "" || ratingStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Nama, Lokasi, dan Rating wajib diisi",
		})
		return
	}

	rating, err := strconv.ParseFloat(ratingStr, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Rating harus berupa angka",
		})
		return
	}

	bioskop := model.Bioskop{
		Nama:   nama,
		Lokasi: lokasi,
		Rating: float32(rating),
	}

	database.DB.Create(&bioskop)

	c.JSON(http.StatusCreated, bioskop)
}

func GetBioskop(c *gin.Context) {
	var bioskops []model.Bioskop

	result := database.DB.Find(&bioskops)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal mengambil data",
		})
		return
	}

	c.JSON(http.StatusOK, bioskops)
}

func UpdateBioskop(c *gin.Context) {
	id := c.Param("id")

	var bioskop model.Bioskop

	if err := database.DB.First(&bioskop, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Data bioskop tidak ditemukan",
		})
		return
	}

	nama := c.PostForm("nama")
	lokasi := c.PostForm("lokasi")
	ratingStr := c.PostForm("rating")

	if nama == "" || lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Nama dan Lokasi tidak boleh kosong",
		})
		return
	}

	var rating float32 = bioskop.Rating
	if ratingStr != "" {
		value, err := strconv.ParseFloat(ratingStr, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Rating harus berupa angka",
			})
			return
		}
		rating = float32(value)
	}

	bioskop.Nama = nama
	bioskop.Lokasi = lokasi
	bioskop.Rating = rating

	database.DB.Save(&bioskop)

	c.JSON(http.StatusOK, bioskop)
}

func DeleteBioskop(c *gin.Context) {

    id := c.Param("id")

    var bioskop model.Bioskop
    if err := database.DB.First(&bioskop, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "Data bioskop tidak ditemukan",
        })
        return
    }

    if err := database.DB.Delete(&bioskop).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Gagal menghapus data bioskop",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Bioskop deleted successfully",
    })
}