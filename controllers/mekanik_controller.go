package controllers

import (
	"net/http"
	"strconv"

	"api_bengkel/config"
	"api_bengkel/models"
	"api_bengkel/utils"

	"github.com/gin-gonic/gin"
)

func GetMekaniks(c *gin.Context) {
	var mekaniks []models.Mekanik
	if err := config.DB.Find(&mekaniks).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to fetch mekaniks")
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "Success", mekaniks)
}

func GetMekanikByID(c *gin.Context) {
	id := c.Param("id")
	var mekanik models.Mekanik
	if err := config.DB.Where("id_mekanik = ?", id).First(&mekanik).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Mekanik not found")
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "Success", mekanik)
}

func CreateMekanik(c *gin.Context) {
	var input models.Mekanik
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.JenisKelamin != "L" && input.JenisKelamin != "P" {
		utils.ResponseError(c, http.StatusBadRequest, "Jenis kelamin harus 'L' atau 'P'")
		return
	}

	if err := config.DB.Create(&input).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to create mekanik")
		return
	}
	utils.ResponseSuccess(c, http.StatusCreated, "Mekanik created", input)
}

func UpdateMekanik(c *gin.Context) {
	id := c.Param("id")
	var mekanik models.Mekanik
	if err := config.DB.Where("id_mekanik = ?", id).First(&mekanik).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Mekanik not found")
		return
	}

	var input models.Mekanik
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.JenisKelamin != "L" && input.JenisKelamin != "P" {
		utils.ResponseError(c, http.StatusBadRequest, "Jenis kelamin harus 'L' atau 'P'")
		return
	}

	mekanik.NamaMekanik = input.NamaMekanik
	mekanik.JenisKelamin = input.JenisKelamin
	mekanik.Alamat = input.Alamat
	mekanik.Telepon = input.Telepon

	if err := config.DB.Save(&mekanik).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to update mekanik")
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "Mekanik updated", mekanik)
}

func DeleteMekanik(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "Invalid mekanik ID")
		return
	}

	if err := config.DB.Delete(&models.Mekanik{}, idInt).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to delete mekanik")
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "Mekanik deleted", nil)
}
