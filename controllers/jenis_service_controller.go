package controllers

import (
	"net/http"
	"strconv"

	"api_bengkel/config"
	"api_bengkel/models"
	"api_bengkel/utils"

	"github.com/gin-gonic/gin"
)

func GetJenisServices(c *gin.Context) {
    var jenisServices []models.JenisService
    if err := config.DB.Find(&jenisServices).Error; err != nil {
        utils.ResponseError(c, http.StatusInternalServerError, "Failed to fetch jenis service")
        return
    }
    utils.ResponseSuccess(c, http.StatusOK, "Success", jenisServices)
}

func GetJenisServiceByID(c *gin.Context) {
    id := c.Param("id")
    var jenisService models.JenisService
    if err := config.DB.Where("id_service = ?", id).First(&jenisService).Error; err != nil {
        utils.ResponseError(c, http.StatusNotFound, "Jenis service not found")
        return
    }
    utils.ResponseSuccess(c, http.StatusOK, "Success", jenisService)
}

func CreateJenisService(c *gin.Context) {
    var input models.JenisService
    if err := c.ShouldBindJSON(&input); err != nil {
        utils.ResponseError(c, http.StatusBadRequest, err.Error())
        return
    }

    if err := config.DB.Create(&input).Error; err != nil {
        utils.ResponseError(c, http.StatusInternalServerError, "Failed to create jenis service")
        return
    }
    utils.ResponseSuccess(c, http.StatusCreated, "Jenis service created", input)
}

func UpdateJenisService(c *gin.Context) {
    id := c.Param("id")
    var jenisService models.JenisService
    if err := config.DB.Where("id_service = ?", id).First(&jenisService).Error; err != nil {
        utils.ResponseError(c, http.StatusNotFound, "Jenis service not found")
        return
    }

    var input models.JenisService
    if err := c.ShouldBindJSON(&input); err != nil {
        utils.ResponseError(c, http.StatusBadRequest, err.Error())
        return
    }

    jenisService.JenisService = input.JenisService

    if err := config.DB.Save(&jenisService).Error; err != nil {
        utils.ResponseError(c, http.StatusInternalServerError, "Failed to update jenis service")
        return
    }

    utils.ResponseSuccess(c, http.StatusOK, "Jenis service updated", jenisService)
}

func DeleteJenisService(c *gin.Context) {
    id := c.Param("id")
    idInt, err := strconv.Atoi(id)
    if err != nil {
        utils.ResponseError(c, http.StatusBadRequest, "Invalid jenis service ID")
        return
    }

    if err := config.DB.Delete(&models.JenisService{}, idInt).Error; err != nil {
        utils.ResponseError(c, http.StatusInternalServerError, "Failed to delete jenis service")
        return
    }

    utils.ResponseSuccess(c, http.StatusOK, "Jenis service deleted", nil)
}
