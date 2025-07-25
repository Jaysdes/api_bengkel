package controllers

import (
	"net/http"
	"strconv"

	"api_bengkel/config"
	"api_bengkel/models"
	"api_bengkel/utils"

	"github.com/gin-gonic/gin"
)

func GetJenisJasas(c *gin.Context) {
	var jenisJasas []models.JenisJasa
	if err := config.DB.Find(&jenisJasas).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to fetch jenis jasa")
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "Success", jenisJasas)
}

func GetJenisJasaByID(c *gin.Context) {
	id := c.Param("id")
	var jenisJasa models.JenisJasa
	if err := config.DB.Where("id_jasa = ?", id).First(&jenisJasa).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Jenis jasa not found")
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "Success", jenisJasa)
}

func CreateJenisJasa(c *gin.Context) {
	var input models.JenisJasa
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := config.DB.Create(&input).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to create jenis jasa")
		return
	}
	utils.ResponseSuccess(c, http.StatusCreated, "Jenis jasa created", input)
}

func UpdateJenisJasa(c *gin.Context) {
	id := c.Param("id")
	var jenisJasa models.JenisJasa
	if err := config.DB.Where("id_jasa = ?", id).First(&jenisJasa).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Jenis jasa not found")
		return
	}

	var input models.JenisJasa
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	jenisJasa.NamaJasa = input.NamaJasa
	jenisJasa.HargaJasa = input.HargaJasa

	if err := config.DB.Save(&jenisJasa).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to update jenis jasa")
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "Jenis jasa updated", jenisJasa)
}

func DeleteJenisJasa(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "Invalid jenis jasa ID")
		return
	}

	if err := config.DB.Delete(&models.JenisJasa{}, idInt).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to delete jenis jasa")
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "Jenis jasa deleted", nil)
}
