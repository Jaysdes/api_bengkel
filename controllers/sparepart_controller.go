package controllers

import (
	"net/http"

	"api_bengkel/config"
	"api_bengkel/models"
	"api_bengkel/utils"

	"github.com/gin-gonic/gin"
)

func GetSparepartByID(c *gin.Context) {
	id := c.Param("id")
	var sparepart models.Sparepart
	if err := config.DB.Where("id_sparepart = ?", id).First(&sparepart).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Sparepart not found")
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "Success", sparepart)
}

// GET /sparepart
func GetSpareparts(c *gin.Context) {
	var spareparts []models.Sparepart
	if err := config.DB.Find(&spareparts).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Gagal mengambil data sparepart")
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "Berhasil mengambil data sparepart", spareparts)
}

// POST /sparepart
func CreateSparepart(c *gin.Context) {
	var input models.Sparepart
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := config.DB.Create(&input).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Gagal menambahkan sparepart")
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "Sparepart berhasil ditambahkan", input)
}

// PUT /sparepart/:id
func UpdateSparepart(c *gin.Context) {
	id := c.Param("id")
	var sparepart models.Sparepart
	if err := config.DB.Where("id_sparepart = ?", id).First(&sparepart).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Sparepart tidak ditemukan")
		return
	}

	if err := c.ShouldBindJSON(&sparepart); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := config.DB.Save(&sparepart).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Gagal mengupdate sparepart")
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "Sparepart berhasil diupdate", sparepart)
}

// DELETE /sparepart/:id
func DeleteSparepart(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Where("id_sparepart = ?", id).Delete(&models.Sparepart{}).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Gagal menghapus sparepart")
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "Sparepart berhasil dihapus", nil)
}

// POST /sparepart/kurangi_stok
func KurangiStokSparepart(c *gin.Context) {
	var input struct {
		IDSparepart string `json:"id_sparepart" binding:"required"`
		Pengurangan int    `json:"pengurangan" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	var sparepart models.Sparepart
	if err := config.DB.Where("id_sparepart = ?", input.IDSparepart).First(&sparepart).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Sparepart tidak ditemukan")
		return
	}

	if sparepart.Stok < input.Pengurangan {
		utils.ResponseError(c, http.StatusBadRequest, "Stok tidak mencukupi")
		return
	}

	sparepart.Stok -= input.Pengurangan

	if err := config.DB.Save(&sparepart).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Gagal mengurangi stok")
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "Stok berhasil dikurangi", sparepart)
}
