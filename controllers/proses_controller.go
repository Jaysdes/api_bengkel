package controllers

import (
	"api_bengkel/config"
	"api_bengkel/models"
	"api_bengkel/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET /proses
func GetAllProses(c *gin.Context) {
	var data []models.Proses
	if err := config.DB.Find(&data).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "List Proses", data)
}

// POST /proses
func CreateProses(c *gin.Context) {
	var input models.Proses
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := config.DB.Create(&input).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseSuccess(c, http.StatusCreated, "Data proses berhasil ditambahkan", input)
}

// PUT /proses/:id
func UpdateProses(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data models.Proses
	if err := config.DB.First(&data, id).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "Proses tidak ditemukan")
		return
	}

	var input models.Proses
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	config.DB.Model(&data).Updates(input)
	utils.ResponseSuccess(c, http.StatusOK, "Data proses berhasil diperbarui", data)
}

// DELETE /proses/:id
func DeleteProses(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.Delete(&models.Proses{}, id).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "Data proses berhasil dihapus", nil)
}
