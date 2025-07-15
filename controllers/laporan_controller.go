package controllers

import (
	"api_bengkel/config"
	"api_bengkel/models"
	"api_bengkel/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET /laporan
func GetAllLaporan(c *gin.Context) {
	var data []models.Laporan
	if err := config.DB.Find(&data).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "List Laporan", data)
}

// POST /laporan
func CreateLaporan(c *gin.Context) {
	var input models.Laporan
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := config.DB.Create(&input).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseSuccess(c, http.StatusCreated, "Data laporan berhasil ditambahkan", input)
}

// PUT /laporan/:id
func UpdateLaporan(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data models.Laporan
	if err := config.DB.First(&data, id).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Laporan tidak ditemukan")
		return
	}

	var input models.Laporan
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	config.DB.Model(&data).Updates(input)
	utils.ResponseSuccess(c, http.StatusOK, "Data laporan berhasil diperbarui", data)
}

// DELETE /laporan/:id
func DeleteLaporan(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.Delete(&models.Laporan{}, id).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "Data laporan berhasil dihapus", nil)
}
