package controllers

import (
	"api_bengkel/config"
	"api_bengkel/models"
	"api_bengkel/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProses(c *gin.Context) {
	var proses []models.Proses
	if err := config.DB.Find(&proses).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to fetch proses")
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "Success", proses)
}

func GetProsesByID(c *gin.Context) {
	id := c.Param("id")
	var proses models.Proses
	if err := config.DB.Where("id_proses = ?", id).First(&proses).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Proses not found")
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "Success", proses)
}

func UpdateProses(c *gin.Context) {
	id := c.Param("id")
	var proses models.Proses
	if err := config.DB.Where("id_proses = ?", id).First(&proses).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Proses not found")
		return
	}

	var input models.Proses
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	proses.Status = input.Status
	proses.Keterangan = input.Keterangan
	proses.WaktuMulai = input.WaktuMulai
	proses.WaktuSelesai = input.WaktuSelesai

	if err := config.DB.Save(&proses).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to update proses")
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "Proses updated", proses)
}

func DeleteProses(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Proses{}, id).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to delete proses")
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "Proses deleted", nil)
}
